package http

import "github.com/google/uuid"

type RegisterData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Account struct {
	ID       uuid.UUID
	Username string
	Password string
}
