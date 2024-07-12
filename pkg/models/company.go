package models

import (
	"errors"
	"fmt"

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

	maxNameLen = 15
	maxDescLen = 3000
)

var (
	invalidName        = "Invalid company name, must be no more than 15 characters"
	invalidDescription = "Invalid description, must be no more than 3000 characters or be empty"
	invalidType        = fmt.Sprintf("Invalid company type, the type must be one of the following: %s, %s, %s, %s", Corporation, NonProfit, Cooperative, SoleProprietorship)
	separator          = "; "
)

type Company struct {
	ID                uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
	Name              string      `gorm:"type:varchar(15);unique;not null" json:"name"`
	Description       string      `gorm:"type:varchar(3000)" json:"description,omitempty"`
	AmountOfEmployees int         `gorm:"not null" json:"amount_of_employees"`
	Registered        bool        `gorm:"not null" json:"registered"`
	Type              CompanyType `gorm:"type:varchar(20);not null" json:"type"`
}

func (c *Company) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

func (c *Company) Validate() error {
	var errStr string
	if !c.validateName() {
		errStr += invalidName
	}

	if !c.validateDescription() {
		if len(errStr) != 0 {
			errStr += separator
		}
		errStr += invalidDescription
	}

	if !c.validateType() {
		if len(errStr) != 0 {
			errStr += separator
		}
		errStr += invalidType
	}

	if len(errStr) != 0 {
		return errors.New(errStr)
	}

	return nil
}

func (c *Company) validateName() bool {
	return len(c.Name) <= maxNameLen
}

func (c *Company) validateDescription() bool {
	return len(c.Description) == 0 || len(c.Description) <= maxDescLen
}

func (c *Company) validateType() bool {
	return c.Type == Corporation || c.Type == NonProfit || c.Type == Cooperative || c.Type == SoleProprietorship
}
