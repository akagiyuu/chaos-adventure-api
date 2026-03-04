package http

import (
	"context"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/lestrrat-go/jwx/v3/jwa"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

const (
	TokenKey string = "token"
)

func (s *Server) RequireToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwt.ParseRequest(r, jwt.WithKey(jwa.RS256(), s.Auth.PublicKey))
		if err != nil {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Err:    err,
				Detail: "Invalid authorization token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), TokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
