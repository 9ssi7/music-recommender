package app

import (
	"github.com/9ssi7/music-recommender/app/command"
	"github.com/9ssi7/music-recommender/app/query"
)

type Commands struct {
	UserCreate command.UserCreateHandler
}

type Queries struct {
	UserGetByEmail query.UserGetByEmailHandler
}

type App struct {
	Commands Commands
	Queries  Queries
}
