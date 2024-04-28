package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/genre"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type GenreDeleteCommand struct {
	Id uuid.UUID
}

type GenreDeleteResult struct {
}

type GenreDeleteHandler cqrs.HandlerFunc[GenreDeleteCommand, *GenreDeleteResult]

func NewGenreDeleteHandler(repo genre.Repo) GenreDeleteHandler {
	return func(ctx context.Context, cmd GenreDeleteCommand) (*GenreDeleteResult, *i18np.Error) {
		err := repo.Delete(ctx, cmd.Id.String())
		if err != nil {
			return nil, err
		}
		return &GenreDeleteResult{}, nil
	}
}
