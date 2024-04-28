package main

import (
	"fmt"

	"github.com/9ssi7/music-recommender/config"
)

func main() {
	cnf := config.Get()
	fmt.Println(cnf)
}
