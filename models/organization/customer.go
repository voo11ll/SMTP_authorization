package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID                 uuid.UUID              `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name               string                 `json:"name"`
	FullName           string                 `json:"fullName"`
	INN                int32                  `json:"inn"`
	KPP                int32                  `json:"kpp"`
	LegalAddress       string                 `json:"legalAddress"`
	Banks              []*CustomerBank        `gorm:"ForeignKey:CustomerID"`
	ContactInfos       []*CustomerContactInfo `gorm:"ForeignKey:CustomerID"`
	Employees          []*CustomerUser        `gorm:"ForeignKey:CustomerID"`
	BusinessUniverseID uuid.UUID              `gorm:"column:business_universe_id"`
}

type CustomerBank struct {
	gorm.Model
	ID                uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name              string    `json:"name"`
	AccountNumber     string    `json:"accountNumber"`
	Bik               string    `json:"bik"`
	CorrAccountNumber string    `json:"corrAccountNumber"`
	IsOpen            bool      `json:"isOpen"`
	CustomerID        uuid.UUID `gorm:"column:customer_id"`
}

type CustomerContactInfo struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ContactTypeID uuid.UUID `gorm:"column:contact_type_id"`
	Value         string    `json:"value"`
	CustomerID    uuid.UUID `gorm:"column:customer_id"`
}
