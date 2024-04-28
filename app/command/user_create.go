package command

import (
	"context"

	"github.com/9ssi7/music-recommender/domains/user"
	"github.com/cilloparch/cillop/cqrs"
	"github.com/cilloparch/cillop/i18np"
	"github.com/google/uuid"
)

type UserCreateCmd struct {
	Email          string
	UserName       string
	FavoriteGenres []string
}

type UserCreateRes struct {
	Dto *user.ListDto
}

type UserCreateHandler cqrs.HandlerFunc[UserCreateCmd, *UserCreateRes]

func NewUserCreateHandler(repo user.Repo) UserCreateHandler {
	return func(ctx context.Context, cmd UserCreateCmd) (*UserCreateRes, *i18np.Error) {
		genres := make([]uuid.UUID, len(cmd.FavoriteGenres))
		for i, genre := range cmd.FavoriteGenres {
			genres[i] = uuid.MustParse(genre)
		}
		dto, err := repo.Create(ctx, user.CreateDto{
			UserName:       cmd.Email,
			Email:          cmd.Email,
			FavoriteGenres: genres,
		})
		if err != nil {
			return nil, err
		}
		return &UserCreateRes{
			Dto: dto,
		}, nil
	}
}
