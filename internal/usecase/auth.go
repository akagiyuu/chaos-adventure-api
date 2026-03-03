package usecase

import (
	"fmt"
	"go/printer"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"

	"github.com/akagiyuu/chaos-adventure-api/internal/config"
	"github.com/akagiyuu/chaos-adventure-api/internal/ports"
)

type Auth struct {
	privkey   jwk.Key
	pubkey    jwk.Key
	expiredIn time.Duration
	repo      ports.Repository
}

func NewAuth(
	cfg *config.Config,
	repo ports.Repository,
) (*Auth, error) {
	privkey, err := jwk.ParseKey(cfg.JWTSecret)
	if err != nil {
		return nil, err
	}

	pubkey, err := jwk.PublicKeyOf(privkey)
	if err != nil {
		return nil, err
	}

	return &Auth{
		privkey:   privkey,
		pubkey:    pubkey,
		expiredIn: time.Duration(cfg.JWTExpiredIn) * time.Hour,
		repo:      repo,
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

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256(), u.privkey))
	if err != nil {
		return nil, err
	}

	return signed, err
}

func (u *Auth) ParseToken(raw []byte) (uuid.UUID, error) {
	token, err := jwt.Parse(raw, jwt.WithKey(jwa.RS256(), u.pubkey))
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
