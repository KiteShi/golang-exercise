package models

import (
	"database/sql"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CompanyType is a custom type for the type of company
type CompanyType string

// Constants for the different types of companies
const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID                uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name              string         `gorm:"type:varchar(15);unique;not null" json:"name"`
	Description       sql.NullString `gorm:"type:varchar(3000)" json:"description,omitempty"`
	AmountOfEmployees int            `gorm:"not null" json:"amount_of_employees"`
	Registered        bool           `gorm:"not null" json:"registered"`
	Type              CompanyType    `gorm:"type:varchar(20);not null" json:"type"`
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
