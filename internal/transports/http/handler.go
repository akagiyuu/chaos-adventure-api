package http

import (
	"net/http"

	"github.com/go-fuego/fuego"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	fuego.Get(f, "/", s.Ping)
}

func (s *Server) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}

func (s *Server) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
