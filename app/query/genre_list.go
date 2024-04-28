package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/genre"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type GenreListQuery struct {
}

type GenreListResult struct {
	Dtos []genre.ListDto
}

type GenreListHandler cqrs.HandlerFunc[GenreListQuery, *GenreListResult]

func NewGenreListHandler(repo genre.Repo) GenreListHandler {
	return func(ctx context.Context, query GenreListQuery) (*GenreListResult, *i18np.Error) {
		dtos, err := repo.List(ctx)
		if err != nil {
			return nil, err
		}
		return &GenreListResult{
			Dtos: dtos,
		}, nil
	}
}
