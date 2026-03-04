package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-fuego/fuego"
)

const (
	authorization string = "Authorization"
	bearer        string = "Bearer "
	TokenKey      string = "token"
)

func (s *Server) RequireToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authorization)
		if authHeader == "" {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization header",
			})
			return
		}

		token, isBearer := strings.CutPrefix(authHeader, bearer)
		if !isBearer {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization token",
			})
			return
		}

		id, err := s.Auth.ParseToken([]byte(token))
		if err != nil {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Err:    err,
				Detail: "Invalid authorization token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), TokenKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
