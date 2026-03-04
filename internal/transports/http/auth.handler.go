package http

import (
	"github.com/akagiyuu/chaos-adventure-api/internal/domain"
	"github.com/go-fuego/fuego"
	"github.com/jinzhu/copier"
)

func (s *Server) Register(c fuego.ContextWithBody[RegisterData]) ([]byte, error) {
	ctx := c.Context()

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	var data domain.RegisterData
	copier.Copy(&data, &body)
	id, err := s.Auth.Register(ctx, data)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Account with given email already existed",
		}
	}

	token, err := s.Auth.CreateToken(id)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return token, nil
}

func (s *Server) Login(c fuego.ContextWithBody[LoginData]) ([]byte, error) {
	ctx := c.Context()

	body, err := c.Body()
	if err != nil {
		return nil, err
	}

	var data domain.LoginData
	copier.Copy(&data, &body)
	id, err := s.Auth.Login(ctx, data)
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Wrong email or password",
		}
	}

	token, err := s.Auth.CreateToken(id)
	if err != nil {
		return nil, fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return token, nil
}
