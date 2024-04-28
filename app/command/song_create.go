package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type SongCreateCommand struct {
	Title  string
	Artist string
}

type SongCreateResult struct {
	Dto *song.ListDto
}

type SongCreateHandler cqrs.HandlerFunc[SongCreateCommand, *SongCreateResult]

func NewSongCreateHandler(repo song.Repo) SongCreateHandler {
	return func(ctx context.Context, cmd SongCreateCommand) (*SongCreateResult, *i18np.Error) {
		dto, err := repo.Create(ctx, song.CreateDto{
			Title:  cmd.Title,
			Artist: cmd.Artist,
		})
		if err != nil {
			return nil, err
		}
		return &SongCreateResult{
			Dto: dto,
		}, nil
	}
}
