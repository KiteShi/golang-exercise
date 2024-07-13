package cmd

import (
	"net/http"

	"github.com/KiteShi/golang-exercise/pkg/handler"
	m "github.com/KiteShi/golang-exercise/pkg/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", handler.LogIn)
	r.Get("/{uuid:[a-zA-Z0-9-_]+}", handler.GetCompany)

	// Protected routes (authentication required)
	r.Group(func(r chi.Router) {
		r.Use(m.AuthMiddleware)

		r.Post("/", handler.CreateCompany)
		r.Patch("/{uuid:[a-zA-Z0-9-_]+}", handler.UpdateCompany)
		r.Delete("/{uuid:[a-zA-Z0-9-_]+}", handler.DeleteCompany)
	})

	return r
}
