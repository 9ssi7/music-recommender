package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type SongDeleteCommand struct {
	Id uuid.UUID
}

type SongDeleteResult struct {
}

type SongDeleteHandler cqrs.HandlerFunc[SongDeleteCommand, *SongDeleteResult]

func NewSongDeleteHandler(repo song.Repo) SongDeleteHandler {
	return func(ctx context.Context, cmd SongDeleteCommand) (*SongDeleteResult, *i18np.Error) {
		err := repo.Delete(ctx, cmd.Id)
		if err != nil {
			return nil, err
		}
		return &SongDeleteResult{}, nil
	}
}
