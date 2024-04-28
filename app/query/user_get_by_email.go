package query

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/user"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
)

type UserGetByEmailQuery struct {
	Email string
}

type UserGetByEmailResult struct {
	Dto *user.ListDto
}

type UserGetByEmailHandler cqrs.HandlerFunc[UserGetByEmailQuery, *UserGetByEmailResult]

func NewUserGetByEmailHandler(repo user.Repo) UserGetByEmailHandler {
	return func(ctx context.Context, query UserGetByEmailQuery) (*UserGetByEmailResult, *i18np.Error) {
		dto, err := repo.GetByEmail(ctx, query.Email)
		if err != nil {
			return nil, err
		}
		return &UserGetByEmailResult{
			Dto: dto,
		}, nil
	}
}
