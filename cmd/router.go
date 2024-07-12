package cmd

import (
	"net/http"

	"github.com/KiteShi/golang-exercise/internal/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func GetRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{uuid:[a-zA-Z0-9-_]+}", api.GetCompany)
	r.Post("/", api.CreateCompany)
	r.Put("/", api.UpdateCompany)
	r.Delete("/{uuid:[a-zA-Z0-9-_]+}", api.DeleteCompany)

	return r
}
