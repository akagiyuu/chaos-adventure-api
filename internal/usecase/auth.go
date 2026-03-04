package usecase

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/akagiyuu/chaos-adventure-api/internal/config"
	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/akagiyuu/chaos-adventure-api/internal/ports"
)

type Auth struct {
	privateKey jwk.Key
	PublicKey  jwk.Key
	expiredIn  time.Duration
	repo       ports.Repository
}

func NewAuth(
	cfg *config.Config,
	repo ports.Repository,
) (Auth, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	privateKey, err := jwk.Import(key)
	if err != nil {
		return Auth{}, err
	}

	publicKey, err := jwk.PublicKeyOf(privateKey)
	if err != nil {
		return Auth{}, err
	}

	return Auth{
		privateKey: privateKey,
		PublicKey:  publicKey,
		expiredIn:  time.Duration(cfg.JWTExpiredIn) * time.Hour,
		repo:       repo,
	}, nil
}

func (u *Auth) CreateToken(id uuid.UUID) ([]byte, error) {
	token, err := jwt.NewBuilder().
		Subject(id.String()).
		Expiration(time.Now().Add(u.expiredIn)).
		Build()
	if err != nil {
		return nil, err
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), u.privateKey))
	if err != nil {
		return nil, err
	}

	return signed, err
}

func (u *Auth) ParseToken(raw []byte) (uuid.UUID, error) {
	token, err := jwt.Parse(raw, jwt.WithKey(jwa.RS256(), u.PublicKey))
	if err != nil {
		return uuid.Nil, err
	}

	sub, exist := token.Subject()
	if !exist {
		return uuid.Nil, fmt.Errorf("Token contain no subject")
	}

	id, err := uuid.Parse(sub)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u *Auth) Register(ctx context.Context, data domain.RegisterData) (uuid.UUID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, err
	}
	data.Password = string(hashedPassword)

	id, err := u.repo.CreateAccount(ctx, data)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u *Auth) Login(ctx context.Context, data domain.LoginData) (uuid.UUID, error) {
	account, err := u.repo.GetAccountByUsername(ctx, data.Username)
	if err != nil {
		return uuid.Nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(data.Password))
	if err != nil {
		return uuid.Nil, err
	}

	return account.ID, nil
}

func (u *Auth) Self(ctx context.Context, id uuid.UUID) (*domain.Account, error) {
	account, err := u.repo.GetAccount(ctx, id)
	if err != nil {
		return nil, err
	}

	return account, nil
}
