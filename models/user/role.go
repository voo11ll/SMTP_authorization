package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name string    `json:"name"`
	User *[]User   `gorm:"ForeignKey:RoleID"`
}
