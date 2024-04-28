package user

import "github.com/google/uuid"

type CreateDto struct {
	Id             uuid.UUID   `json:"-" neo4j:"id"`
	UserName       string      `json:"userName" neo4j:"userName"`
	Email          string      `json:"email" neo4j:"email"`
	FavoriteGenres []uuid.UUID `json:"favoriteGenres" neo4j:"favoriteGenres"`
}

type ListDto struct {
	Id       uuid.UUID `json:"id" neo4j:"id"`
	UserName string    `json:"userName" neo4j:"userName"`
	Email    string    `json:"email" neo4j:"email"`
}

func (dto *CreateDto) Build() map[string]interface{} {
	dto.Id = uuid.New()
	return map[string]interface{}{
		"id":             dto.Id.String(),
		"userName":       dto.UserName,
		"email":          dto.Email,
		"favoriteGenres": dto.FavoriteGenres,
	}
}
