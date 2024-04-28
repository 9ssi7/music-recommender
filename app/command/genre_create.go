package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/genre"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type GenreCreateCommand struct {
	Name string
}

type GenreCreateResult struct {
	Dto *genre.ListDto
}

type GenreCreateHandler cqrs.HandlerFunc[GenreCreateCommand, *GenreCreateResult]

func NewGenreCreateHandler(repo genre.Repo) GenreCreateHandler {
	return func(ctx context.Context, cmd GenreCreateCommand) (*GenreCreateResult, *i18np.Error) {
		dto, err := repo.Create(ctx, genre.CreateDto{
			Name: cmd.Name,
		})
		if err != nil {
			return nil, err
		}
		return &GenreCreateResult{
			Dto: dto,
		}, nil
	}
}
