package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactType struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name   string    `json:"name"`
	Notion string    `json:"notion"`
}
