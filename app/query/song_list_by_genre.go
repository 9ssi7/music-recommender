package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type SongListByGenreQuery struct {
	GenreId uuid.UUID
}

type SongListByGenreResult struct {
	Dtos []song.ListDto
}

type SongListByGenreHandler cqrs.HandlerFunc[SongListByGenreQuery, *SongListByGenreResult]

func NewSongListByGenreHandler(repo song.Repo) SongListByGenreHandler {
	return func(ctx context.Context, query SongListByGenreQuery) (*SongListByGenreResult, *i18np.Error) {
		dtos, err := repo.ListByGenre(ctx, query.GenreId)
		if err != nil {
			return nil, err
		}
		return &SongListByGenreResult{
			Dtos: dtos,
		}, nil
	}
}
