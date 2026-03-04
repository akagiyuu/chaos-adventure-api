package http

import (
	"net/http"

	"github.com/akagiyuu/chaos-adventure-api/internal/config"
	"github.com/akagiyuu/chaos-adventure-api/internal/usecase"
	"github.com/go-fuego/fuego"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Handler struct {
	Config *config.Config
	Auth   usecase.Auth
}

func (h *Handler) RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/", h.Ping)
}

func (h *Handler) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}

func (h *Handler) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
