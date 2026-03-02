package repo

import (
	"context"
	"fmt"

	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/akagiyuu/chaos-adventure-api/internal/ports"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	accounts map[uuid.UUID]domain.Account
}

var _ ports.Repository = &MemoryRepository{}

func NewMemoryRepository() MemoryRepository {
	return MemoryRepository{
		accounts: make(map[uuid.UUID]domain.Account, 0),
	}
}

func (r *MemoryRepository) CreateAccount(ctx context.Context, username string, password string) (uuid.UUID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	account := domain.Account{
		ID:       id,
		Username: username,
		Password: password,
	}

	r.accounts[id] = account

	return id, nil
}

func (r *MemoryRepository) GetAccount(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	account, exist := r.accounts[id]
	if !exist {
		return nil, fmt.Errorf("Account not found")
	}

	return &account, nil
}
