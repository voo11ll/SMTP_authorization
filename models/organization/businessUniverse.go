package models

import (
	user "auth/auth_back/models/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessUniverse struct {
	gorm.Model
	ID       uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string       `json:"name"`
	Domain   string       `json:"domain"`
	User     []*user.User `gorm:"ForeignKey:BusinessUniverseID"`
	Company  []*Company   `gorm:"ForeignKey:BusinessUniverseID"`
	Customer []*Customer  `gorm:"ForeignKey:BusinessUniverseID"`
}
