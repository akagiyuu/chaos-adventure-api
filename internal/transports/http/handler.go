package http

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	fuego.Get(f, "/", s.Ping)

	auth := fuego.Group(f, "/auth")
	fuego.Post(auth, "/register", s.Register)
	fuego.Post(auth, "/login", s.Login)
	fuego.Get(auth, "/self", s.Self,
		option.Middleware(s.RequireToken),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)

	record := fuego.Group(f, "/record")
	fuego.Post(record, "/", s.CreateRecord,
		option.Middleware(s.RequireToken),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Get(record, "/", s.GetAllRecord)
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
