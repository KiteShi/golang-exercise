package cmd

import (
	"net/http"

	"github.com/KiteShi/golang-exercise/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{uuid:[a-zA-Z0-9-_]+}", api.GetCompany)
	r.Post("/", api.CreateCompany)
	r.Put("/{uuid:[a-zA-Z0-9-_]+}", api.UpdateCompany)
	r.Delete("/{uuid:[a-zA-Z0-9-_]+}", api.DeleteCompany)

	return r
}
