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
	return func(ctx context.Context, Command GenreDeleteCommand) (*GenreDeleteResult, *i18np.Error) {
		err := repo.Delete(ctx, Command.Id.String())
		if err != nil {
			return nil, err
		}
		return &GenreDeleteResult{}, nil
	}
}
 