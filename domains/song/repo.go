package song

import (
	"context"

	"github.com/9ssi7/music-recommender/pkg/cypher"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repo interface {
	List(ctx context.Context) ([]ListDto, *i18np.Error)
	ListByGenre(ctx context.Context, genreId uuid.UUID) ([]ListDto, *i18np.Error)
	ListUserRecommendation(ctx context.Context, userId uuid.UUID) ([]ListDto, *i18np.Error)
	View(ctx context.Context, id uuid.UUID) (*ListDto, *i18np.Error)
	Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error)
	Delete(ctx context.Context, id uuid.UUID) *i18np.Error
	MarkListened(ctx context.Context, userId, songId uuid.UUID) *i18np.Error
}

type repo struct {
	driver neo4j.DriverWithContext
}

func NewRepo(driver neo4j.DriverWithContext) Repo {
	return &repo{driver}
}

func (r *repo) List(ctx context.Context) ([]ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (s:Song)
		RETURN s.id, s.title, s.artist
	`
	record, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	var listDto []ListDto
	for record.Next(ctx) {
		var dto ListDto
		if err := cypher.Parse(record.Record(), "s", &dto); err != nil {
			return nil, i18np.NewError(Messages.ParseFailed)
		}
		listDto = append(listDto, dto)
	}
	return listDto, nil
}

func (r *repo) ListByGenre(ctx context.Context, genreId uuid.UUID) ([]ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (g:Genre {id: $genreId})-[:HAS]->(s:Song)
		RETURN s.id, s.title, s.artist
	`
	args := map[string]interface{}{"genreId": genreId}
	record, err := session.Run(ctx, query, args)
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	var listDto []ListDto
	for record.Next(ctx) {
		var dto ListDto
		if err := cypher.Parse(record.Record(), "s", &dto); err != nil {
			return nil, i18np.NewError(Messages.ParseFailed)
		}
		listDto = append(listDto, dto)
	}
	return listDto, nil
}

func (r *repo) ListUserRecommendation(ctx context.Context, userId uuid.UUID) ([]ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (u:User {id: "user_id"})-[:LISTENED_TO]->(s:Song)<-[:HAS_GENRE]-(g:Genre)<-[:HAS_GENRE]-(recommendedSong:Song)
		WHERE NOT (u)-[:LISTENED_TO]->(recommendedSong) AND recommendedSong <> s
		RETURN recommendedSong.id, recommendedSong.title, recommendedSong.artist, COUNT(DISTINCT g) AS commonGenres
		ORDER BY commonGenres DESC
		LIMIT 10
	`
	args := map[string]interface{}{"userId": userId}
	record, err := session.Run(ctx, query, args)
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	var listDto []ListDto
	for record.Next(ctx) {
		var dto ListDto
		if err := cypher.Parse(record.Record(), "recommendedSong", &dto); err != nil {
			return nil, i18np.NewError(Messages.ParseFailed)
		}
		listDto = append(listDto, dto)
	}
	return listDto, nil
}

func (r *repo) View(ctx context.Context, id uuid.UUID) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (s:Song {id: $id})
		RETURN s.id, s.title, s.artist
	`
	args := map[string]interface{}{"id": id}
	record, err := session.Run(ctx, query, args)
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	if !record.Next(ctx) {
		return nil, nil
	}
	var dto ListDto
	if err := cypher.Parse(record.Record(), "s", &dto); err != nil {
		return nil, i18np.NewError(Messages.ParseFailed)
	}
	return &dto, nil
}

func (r *repo) Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	query := `
		CREATE (s:Song {id: $id, title: $title, artist: $artist})
		RETURN s.id, s.title, s.artist
	`
	record, err := session.Run(ctx, query, dto.Build())
	if err != nil {
		return nil, i18np.NewError(Messages.CreateFailed)
	}
	if !record.Next(ctx) {
		return nil, nil
	}
	var listDto ListDto
	if err := cypher.Parse(record.Record(), "s", &listDto); err != nil {
		return nil, i18np.NewError(Messages.ParseFailed)
	}
	return &listDto, nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) *i18np.Error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	query := `
		MATCH (s:Song {id: $id})
		DETACH DELETE s
	`
	args := map[string]interface{}{"id": id}
	_, err := session.Run(ctx, query, args)
	if err != nil {
		return i18np.NewError(Messages.DeleteFailed)
	}
	return nil
}

func (r *repo) MarkListened(ctx context.Context, userId, songId uuid.UUID) *i18np.Error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	query := `
		MATCH (u:User {id: $userId})
		MATCH (s:Song {id: $songId})
		MERGE (u)-[r:LISTENED]->(s)
		SET r.listenedAt = datetime()
	`
	args := map[string]interface{}{"userId": userId, "songId": songId}
	_, err := session.Run(ctx, query, args)
	if err != nil {
		return i18np.NewError(Messages.MarkListenedFailed)
	}
	return nil
}