package genre

import "github.com/google/uuid"

type ListDto struct {
	Id   uuid.UUID `json:"id" neo4j:"id"`
	Name string    `json:"name" neo4j:"name"`
}

type CreateDto struct {
	Name string `json:"name" neo4j:"name"`
}

func (dto *CreateDto) Build() map[string]interface{} {
	return map[string]interface{}{
		"id":   uuid.New(),
		"name": dto.Name,
	}
}
