package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jinzhu/copier"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"

	"github.com/akagiyuu/chaos-adventure-api/internal/adapters/repo/database"
	"github.com/akagiyuu/chaos-adventure-api/internal/config"
	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/akagiyuu/chaos-adventure-api/internal/ports"
)

type Database struct {
	pool *pgxpool.Pool
}

var _ ports.Repository = Database{}

func NewDatabase(cfg *config.Config) (Database, error) {
	pgxConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return Database{}, err
	}

	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		return Database{}, nil
	}

	return Database{pool}, nil
}

func (db Database) CreateAccount(ctx context.Context, data domain.RegisterData) (uuid.UUID, error) {
	inner := database.New(db.pool)

	var params database.CreateAccountParams
	copier.Copy(&params, &data)

	return inner.CreateAccount(ctx, params)
}

func (db Database) GetAccount(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	inner := database.New(db.pool)

	raw, err := inner.GetAccount(ctx, id)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var account domain.Account
	copier.Copy(&account, &raw)

	return &account, nil
}

func (db Database) GetAccountByUsername(ctx context.Context, username string) (*domain.Account, error) {
	inner := database.New(db.pool)

	raw, err := inner.GetAccountByUsername(ctx, username)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	var account domain.Account
	copier.Copy(&account, &raw)

	return &account, nil
}
