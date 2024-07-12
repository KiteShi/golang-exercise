package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	//TODO: add handler logic, remove stub line
	w.Write([]byte(id))
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	//TODO: add handler logic
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	//TODO: add handler logic
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "uuid")

	//TODO: add handler logic, remove stub line
	w.Write([]byte(id))
}
