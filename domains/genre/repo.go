package genre

import (
	"context"

	"github.com/9ssi7/music-recommender/pkg/cypher"
	"github.com/cilloparch/cillop/i18np"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Repo interface {
	List(ctx context.Context) ([]ListDto, *i18np.Error)
	View(ctx context.Context, id string) (*ListDto, *i18np.Error)
	Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error)
	Delete(ctx context.Context, id string) *i18np.Error
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
		MATCH (g:Genre)
		RETURN g.id, g.name
	`
	record, err := session.Run(ctx, query, nil)
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	var listDto []ListDto
	for record.Next(ctx) {
		var dto ListDto
		if err := cypher.Parse(record.Record(), "g", &dto); err != nil {
			return nil, i18np.NewError(Messages.ParseFailed)
		}
		listDto = append(listDto, dto)
	}
	return listDto, nil
}

func (r *repo) View(ctx context.Context, id string) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close(ctx)
	query := `
		MATCH (g:Genre {id: $id})
		RETURN g.id, g.name
	`
	record, err := session.Run(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, i18np.NewError(Messages.FetchFailed)
	}
	if !record.Next(ctx) {
		return nil, nil
	}
	var dto ListDto
	if err := cypher.Parse(record.Record(), "g", &dto); err != nil {
		return nil, i18np.NewError(Messages.ParseFailed)
	}
	return &dto, nil
}

func (r *repo) Create(ctx context.Context, dto CreateDto) (*ListDto, *i18np.Error) {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	query := `
		CREATE (g:Genre {id: $id, name: $name})
		RETURN g.id, g.name
	`
	record, err := session.Run(ctx, query, dto.Build())
	if err != nil {
		return nil, i18np.NewError(Messages.CreateFailed)
	}
	if !record.Next(ctx) {
		return nil, nil
	}
	var listDto ListDto
	if err := cypher.Parse(record.Record(), "g", &listDto); err != nil {
		return nil, i18np.NewError(Messages.ParseFailed)
	}
	return &listDto, nil
}

func (r *repo) Delete(ctx context.Context, id string) *i18np.Error {
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	query := `
		MATCH (g:Genre {id: $id})
		DETACH DELETE g
	`
	_, err := session.Run(ctx, query, map[string]interface{}{"id": id})
	if err != nil {
		return i18np.NewError(Messages.DeleteFailed)
	}
	return nil
}
