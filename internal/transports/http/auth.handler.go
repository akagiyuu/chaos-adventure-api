package http

import (
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"

	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
)

func (s *Server) Register(c fuego.ContextWithBody[RegisterData]) (string, error) {
	ctx := c.Context()

	body, err := c.Body()
	if err != nil {
		return "", err
	}

	var data domain.RegisterData
	copier.Copy(&data, &body)
	id, err := s.Auth.Register(ctx, data)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Account with given email already existed",
		}
	}

	token, err := s.Auth.CreateToken(id)
	if err != nil {
		return "", fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return string(token), nil
}

func (s *Server) Login(c fuego.ContextWithBody[LoginData]) (string, error) {
	ctx := c.Context()

	body, err := c.Body()
	if err != nil {
		return "", err
	}

	var data domain.LoginData
	copier.Copy(&data, &body)
	id, err := s.Auth.Login(ctx, data)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Wrong email or password",
		}
	}

	token, err := s.Auth.CreateToken(id)
	if err != nil {
		return "", fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return string(token), nil
}

func (s *Server) Self(c fuego.ContextNoBody) (*Account, error) {
	ctx := c.Context()

	id := c.Value(TokenKey).(uuid.UUID)

	raw, err := s.Auth.Self(ctx, id)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid user",
		}
	}

	var account Account
	copier.Copy(&account, raw)

	return &account, nil
}
