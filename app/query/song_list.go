package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type SongListQuery struct {
}

type SongListResult struct {
	Dtos []song.ListDto
}

type SongListHandler cqrs.HandlerFunc[SongListQuery, *SongListResult]

func NewSongListHandler(repo song.Repo) SongListHandler {
	return func(ctx context.Context, query SongListQuery) (*SongListResult, *i18np.Error) {
		dtos, err := repo.List(ctx)
		if err != nil {
			return nil, err
		}
		return &SongListResult{
			Dtos: dtos,
		}, nil
	}
}
