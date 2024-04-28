package main

import (
	"github.com/9ssi7/music-recommender/config"
	"github.com/9ssi7/music-recommender/pkg/clients/neo4j"
	"github.com/9ssi7/music-recommender/server/graph"
	"github.com/9ssi7/music-recommender/service"
)

func main() {
	cnf := config.Get()
	driver := neo4j.Connect(cnf.Neo4j)
	app := service.NewApp(service.Config{
		Driver: driver,
	})
	graphSrv := graph.NewServer(cnf.Graphql, app)
	graphSrv.Listen()
}
