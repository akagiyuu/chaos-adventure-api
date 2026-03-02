package ports

import (
	"context"

	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	CreateAccount(ctx context.Context, username string, password string) (uuid.UUID, error)
	GetAccount(ctx context.Context, id uuid.UUID) (*domain.Account, error)
}
