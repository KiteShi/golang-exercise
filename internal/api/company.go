package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KiteShi/golang-exercise/pkg/db"
	"github.com/KiteShi/golang-exercise/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetCompany(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID provided", http.StatusBadRequest)
		log.Printf("Error parsing company ID: %v", err)
		return
	}

	company, err := db.GetCompanyByID(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			http.Error(w, "Company not found", http.StatusNotFound)
			log.Printf("Company not found: %v", err)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error fetching company: %v", err)
		}
		return
	}

	encodeResponseToJSON(w, company)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding request payload: %v", err)
		return
	}

	if err := company.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Invalid request payload: %v", err)
		return
	}

	if err := db.DB.Create(&company).Error; err != nil {
		http.Error(w, "Failed to create company", http.StatusInternalServerError)
		log.Printf("Error creating company: %v", err)
		return
	}

	encodeResponseToJSON(w, company)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID provided", http.StatusBadRequest)
		log.Printf("Error parsing company ID: %v", err)
		return
	}

	var company models.Company
	if err := db.DB.First(&company, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Company not found", http.StatusNotFound)
			log.Printf("Company not found: %v", err)
		} else {
			http.Error(w, "Failed to fetch company", http.StatusInternalServerError)
			log.Printf("Error fetching company: %v", err)
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error decoding request payload: %v", err)
		return
	}

	if err := company.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Invalid request payload: %v", err)
		return
	}

	if err := db.DB.Save(&company).Error; err != nil {
		http.Error(w, "Failed to update company", http.StatusInternalServerError)
		return
	}

	encodeResponseToJSON(w, company)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID provided", http.StatusBadRequest)
		log.Printf("Error parsing company ID: %v", err)
		return
	}

	if err := db.DB.Delete(&models.Company{}, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Company not found", http.StatusNotFound)
			log.Printf("Company not found: %v", err)
		} else {
			http.Error(w, "Failed to delete company", http.StatusInternalServerError)
			log.Printf("Error deleting company: %v", err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func encodeResponseToJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)
		return
	}
}
