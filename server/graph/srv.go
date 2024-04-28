package graph

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/9ssi7/music-recommender/app"
	"github.com/9ssi7/music-recommender/config"
	"github.com/cilloparch/cillop/server"
)

type srv struct {
	app app.App
	cnf config.Graphql
}

func NewServer(cnf config.Graphql, app app.App) server.Server {
	return &srv{
		app: app,
		cnf: cnf,
	}
}

func (s *srv) Listen() error {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		app: &s.app,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to %s://%s:%v/ for GraphQL playground", s.cnf.Protocol, s.cnf.Host, s.cnf.Port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%v", s.cnf.Port), nil))

	return nil
}
