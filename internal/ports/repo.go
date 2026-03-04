package ports

import (
	"context"

	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	CreateAccount(ctx context.Context, data domain.RegisterData) (uuid.UUID, error)
	GetAccount(ctx context.Context, id uuid.UUID) (*domain.Account, error)
	GetAccountByUsername(ctx context.Context, username string) (*domain.Account, error)

	CreateRank(ctx context.Context, id uuid.UUID, time float32) error
	GetAllRank(ctx context.Context) ([]domain.Rank, error)
}
