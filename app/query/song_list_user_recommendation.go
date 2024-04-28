package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type SongListUserRecommendationQuery struct {
	UserId uuid.UUID
}

type SongListUserRecommendationResult struct {
	Dtos []song.ListDto
}

type SongListUserRecommendationHandler cqrs.HandlerFunc[SongListUserRecommendationQuery, *SongListUserRecommendationResult]

func NewSongListUserRecommendationHandler(repo song.Repo) SongListUserRecommendationHandler {
	return func(ctx context.Context, query SongListUserRecommendationQuery) (*SongListUserRecommendationResult, *i18np.Error) {
		dtos, err := repo.ListUserRecommendation(ctx, query.UserId)
		if err != nil {
			return nil, err
		}
		return &SongListUserRecommendationResult{
			Dtos: dtos,
		}, nil
	}
}
