package user

import (
	"context"

	"github.com/9ssi7/music-recommender/pkg/cypher"
	"github.com/cilloparch/cillop/i18np"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repo interface {
	GetByEmail(ctx context.Context, email string) (*ListDto, *i18np.Error)
	Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error)
}

type repo struct {
	driver neo4j.DriverWithContext
}

func NewRepo(driver neo4j.DriverWithContext) Repo {
	return &repo{driver}
}

func (r *repo) Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	already, _err := r.GetByEmail(ctx, dto.Email)
	if _err != nil {
		return nil, _err
	}
	if already != nil {
		return nil, i18np.NewError(Messages.EmailAlreadyExists)
	}
	query := `
		CREATE (u:User {id: $id, userName: $userName, email: $email, favoriteGenres: $favoriteGenres})
		WITH u
		UNWIND $favoriteGenres AS genreId
		MATCH (g:Genre {id: genreId})
		CREATE (u)-[:FAVORITE_GENRE]->(g)
		RETURN u.id, u.userName, u.email
	`
	_, err := session.Run(ctx, query, dto.Build())
	if err != nil {
		return nil, i18np.NewError(Messages.CreateFailed)
	}
	return &ListDto{
		Id:       dto.Id,
		UserName: dto.UserName,
		Email:    dto.Email,
	}, nil
}

func (r *repo) GetByEmail(ctx context.Context, email string) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (u:User {email: $email})
		RETURN u.id, u.userName, u.email
	`
	record, err := session.Run(ctx, query, map[string]interface{}{"email": email})
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	if !record.Next(ctx) {
		return nil, nil
	}
	var listDto ListDto
	if err := cypher.Parse(record.Record(), "u", &listDto); err != nil {
		return nil, i18np.NewError(Messages.ParseFailed)
	}
	return &listDto, nil
}
