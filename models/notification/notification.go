package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MailConfrimationLinks struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	HashKey   string    `json:"hashKey"`
	SendTry   int       `json:"sendTry"`
	Confirmed bool      `json:"confirmed" gorm:"default:false"`
	UserId    uuid.UUID `gorm:"column:user_id"`
}
