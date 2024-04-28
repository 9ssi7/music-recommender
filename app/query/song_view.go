package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type SongViewQuery struct {
	Id uuid.UUID
}

type SongViewResult struct {
	Dto *song.ListDto
}

type SongViewHandler cqrs.HandlerFunc[SongViewQuery, *SongViewResult]

func NewSongViewHandler(repo song.Repo) SongViewHandler {
	return func(ctx context.Context, query SongViewQuery) (*SongViewResult, *i18np.Error) {
		dto, err := repo.View(ctx, query.Id)
		if err != nil {
			return nil, err
		}
		return &SongViewResult{
			Dto: dto,
		}, nil
	}
}
