package config

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Neo4j struct {
	Uri      string `yaml:"uri"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Rest struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type Graphql struct {
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	Protocol string `yaml:"protocol"`
}

type App struct {
	Neo4j   Neo4j   `yaml:"neo4j"`
	Rest    Rest    `yaml:"rest"`
	Graphql Graphql `yaml:"graphql"`
}

var configs *App

func Get() *App {
	if configs != nil {
		return configs
	}
	filename, _ := filepath.Abs("./config.yml")
	cleanedDst := filepath.Clean(filename)
	yamlFile, _ := os.ReadFile(cleanedDst)
	err := yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		log.Fatal("error loading config.yml ", err)
	}
	return configs
}
