package app

import (
	"github.com/9ssi7/music-recommender/app/command"
	"github.com/9ssi7/music-recommender/app/query"
)

type Commands struct {
	UserCreate command.UserCreateHandler

	GenreCreate command.GenreCreateHandler
	GenreDelete command.GenreDeleteHandler
}

type Queries struct {
	UserGetByEmail query.UserGetByEmailHandler

	GenreList query.GenreListHandler
	GenreView query.GenreViewHandler
}

type App struct {
	Commands Commands
	Queries  Queries
}
