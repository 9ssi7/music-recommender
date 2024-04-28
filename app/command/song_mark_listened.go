package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type SongMarkListenedCommand struct {
	Id     uuid.UUID
	UserId uuid.UUID
}

type SongMarkListenedResult struct {
}

type SongMarkListenedHandler cqrs.HandlerFunc[SongMarkListenedCommand, *SongMarkListenedResult]

func NewSongMarkListenedHandler(repo song.Repo) SongMarkListenedHandler {
	return func(ctx context.Context, cmd SongMarkListenedCommand) (*SongMarkListenedResult, *i18np.Error) {
		err := repo.MarkListened(ctx, cmd.Id, cmd.UserId)
		if err != nil {
			return nil, err
		}
		return &SongMarkListenedResult{}, nil
	}
}
