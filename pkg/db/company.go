package db

import (
	"github.com/KiteShi/golang-exercise/pkg/models"
	"github.com/google/uuid"
)

func CreateCompany(company *models.Company) error {
	result := DB.Create(company)
	return result.Error
}

func GetCompanyByID(id uuid.UUID) (*models.Company, error) {
	var company models.Company
	result := DB.First(&company, id)
	return &company, result.Error
}

func UpdateCompany(company *models.Company) error {
	result := DB.Save(company)
	return result.Error
}

func DeleteCompanyByID(id uint) error {
	result := DB.Delete(&models.Company{}, id)
	return result.Error
}
