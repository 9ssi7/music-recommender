package song

import "github.com/google/uuid"

type ListDto struct {
	Id     uuid.UUID `json:"id" neo4j:"id"`
	Title  string    `json:"title" neo4j:"title"`
	Artist string    `json:"artist" neo4j:"artist"`
}

type CreateDto struct {
	Title  string `json:"title" neo4j:"title"`
	Artist string `json:"artist" neo4j:"artist"`
}

func (dto *CreateDto) Build() map[string]interface{} {
	return map[string]interface{}{
		"id":     uuid.New(),
		"title":  dto.Title,
		"artist": dto.Artist,
	}
}
