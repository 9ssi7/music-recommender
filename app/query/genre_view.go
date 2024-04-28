package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/genre"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type GenreViewQuery struct {
	Id uuid.UUID
}

type GenreViewResult struct {
	Dto *genre.ListDto
}

type GenreViewHandler cqrs.HandlerFunc[GenreViewQuery, *GenreViewResult]

func NewGenreViewHandler(repo genre.Repo) GenreViewHandler {
	return func(ctx context.Context, query GenreViewQuery) (*GenreViewResult, *i18np.Error) {
		dto, err := repo.View(ctx, query.Id.String())
		if err != nil {
			return nil, err
		}
		return &GenreViewResult{
			Dto: dto,
		}, nil
	}
}
