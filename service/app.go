package service

import (
	"github.com/9ssi7/music-recommender/app"
	"github.com/9ssi7/music-recommender/app/command"
	"github.com/9ssi7/music-recommender/app/query"
	"github.com/9ssi7/music-recommender/domains/genre"
	"github.com/9ssi7/music-recommender/domains/song"
	"github.com/9ssi7/music-recommender/domains/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Config struct {
	Driver neo4j.DriverWithContext
}

func NewApp(cnf Config) app.App {
	userRepo := user.NewRepo(cnf.Driver)
	genreRepo := genre.NewRepo(cnf.Driver)
	songRepo := song.NewRepo(cnf.Driver)

	return app.App{
		Commands: app.Commands{
			UserCreate:       command.NewUserCreateHandler(userRepo),
			GenreCreate:      command.NewGenreCreateHandler(genreRepo),
			GenreDelete:      command.NewGenreDeleteHandler(genreRepo),
			SongCreate:       command.NewSongCreateHandler(songRepo),
			SongDelete:       command.NewSongDeleteHandler(songRepo),
			SongMarkListened: command.NewSongMarkListenedHandler(songRepo),
		},
		Queries: app.Queries{
			UserGetByEmail:             query.NewUserGetByEmailHandler(userRepo),
			GenreList:                  query.NewGenreListHandler(genreRepo),
			GenreView:                  query.NewGenreViewHandler(genreRepo),
			SongList:                   query.NewSongListHandler(songRepo),
			SongView:                   query.NewSongViewHandler(songRepo),
			SongListByGenre:            query.NewSongListByGenreHandler(songRepo),
			SongListUserRecommendation: query.NewSongListUserRecommendationHandler(songRepo),
		},
	}
}
