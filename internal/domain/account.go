package domain

import "github.com/google/uuid"

type Account struct {
	ID       uuid.UUID
	Username string
	Password string
}
