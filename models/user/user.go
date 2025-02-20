package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	SecondName string    `json:"secondName"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Avatar     string    `json:"avatar"`
	RoleID     uuid.UUID `gorm:"column:role_id"`
}
