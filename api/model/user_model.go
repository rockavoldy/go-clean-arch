package model

import "go-clean-arch/entity"

type GetUserResponse struct {
	ID    entity.ID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateUserResponse struct {
	ID    entity.ID `json:"id"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
}
