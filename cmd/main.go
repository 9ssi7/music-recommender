package main

import (
	"fmt"

	"github.com/9ssi7/music-recommender/config"
	"github.com/9ssi7/music-recommender/pkg/clients/neo4j"
)

func main() {
	cnf := config.Get()
	driver := neo4j.Connect(cnf.Neo4j)
	fmt.Println(driver.Target().Scheme)
}
