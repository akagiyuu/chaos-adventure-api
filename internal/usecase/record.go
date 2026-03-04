package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/akagiyuu/chaos-adventure-api/internal/ports"
)

type Record struct {
	Repo ports.Repository
}

func (u *Record) CreateRecord(ctx context.Context, accountID uuid.UUID, time float32) error {
	return u.Repo.CreateRecord(ctx, accountID, time)
}

func (u *Record) GetAllRecord(ctx context.Context) ([]domain.Record, error) {
	return u.Repo.GetAllRecord(ctx)
}
