package genre

import "github.com/google/uuid"

type ListDto struct {
	Id   uuid.UUID `json:"id" neo4j:"id"`
	Name string    `json:"name" neo4j:"name"`
}

type CreateDto struct {
	Id   uuid.UUID `json:"-" neo4j:"id"`
	Name string    `json:"name" neo4j:"name"`
}

func (dto *CreateDto) Build() map[string]interface{} {
	dto.Id = uuid.New()
	return map[string]interface{}{
		"id":   dto.Id.String(),
		"name": dto.Name,
	}
}
