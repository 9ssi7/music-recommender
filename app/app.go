package app

import (
	"github.com/9ssi7/music-recommender/app/command"
	"github.com/9ssi7/music-recommender/app/query"
)

type Commands struct {
	UserCreate command.UserCreateHandler

	GenreCreate command.GenreCreateHandler
	GenreDelete command.GenreDeleteHandler

	SongCreate       command.SongCreateHandler
	SongDelete       command.SongDeleteHandler
	SongMarkListened command.SongMarkListenedHandler
}

type Queries struct {
	UserGetByEmail query.UserGetByEmailHandler

	GenreList query.GenreListHandler
	GenreView query.GenreViewHandler

	SongList                   query.SongListHandler
	SongView                   query.SongViewHandler
	SongListByGenre            query.SongListByGenreHandler
	SongListUserRecommendation query.SongListUserRecommendationHandler
}

type App struct {
	Commands Commands
	Queries  Queries
}
