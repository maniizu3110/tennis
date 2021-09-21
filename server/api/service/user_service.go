package service

import "server/api/sqlc"

type userResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func newUserResponse(user sqlc.User) *userResponse {
	return &userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
	}
}
