package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID                 uuid.UUID             `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name               string                `json:"name"`
	FullName           string                `json:"fullName"`
	INN                int32                 `json:"inn"`
	KPP                int32                 `json:"kpp"`
	LegalAddress       string                `json:"legalAddress"`
	Banks              []*CompanyBank        `gorm:"ForeignKey:CompanyID"`
	ContactInfos       []*CompanyContactInfo `gorm:"ForeignKey:CompanyID"`
	BusinessUniverseID uuid.UUID             `gorm:"column:business_universe_id"`
}

type CompanyBank struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name              string    `json:"name"`
	AccountNumber     string    `json:"accountNumber"`
	Bik               string    `json:"bik"`
	CorrAccountNumber string    `json:"corrAccountNumber"`
	IsOpen            bool      `json:"isOpen"`
	CompanyID         uuid.UUID `gorm:"column:company_id"`
}

type CompanyContactInfo struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ContactTypeID uuid.UUID `gorm:"column:contact_type_id"`
	Value         string    `json:"value"`
	CompanyID     uuid.UUID `gorm:"column:company_id"`
}
