package user

import "github.com/google/uuid"

type CreateDto struct {
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
	return map[string]interface{}{
		"id":             uuid.New(),
		"userName":       dto.UserName,
		"email":          dto.Email,
		"favoriteGenres": dto.FavoriteGenres,
	}
}
