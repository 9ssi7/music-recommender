package song

import "github.com/google/uuid"

type ListDto struct {
	Id     uuid.UUID `json:"id" neo4j:"id"`
	Title  string    `json:"title" neo4j:"title"`
	Artist string    `json:"artist" neo4j:"artist"`
}

type CreateDto struct {
	Id      uuid.UUID `json:"-" neo4j:"id"`
	Title   string    `json:"title" neo4j:"title"`
	Artist  string    `json:"artist" neo4j:"artist"`
	GenreId uuid.UUID `json:"genreId" neo4j:"genreId"`
}

func (dto *CreateDto) Build() map[string]interface{} {
	dto.Id = uuid.New()
	return map[string]interface{}{
		"id":      dto.Id.String(),
		"title":   dto.Title,
		"artist":  dto.Artist,
		"genreId": dto.GenreId.String(),
	}
}
