package domain

import "github.com/google/uuid"

type Rank struct {
	AccountID uuid.UUID
	Time      float32
}
