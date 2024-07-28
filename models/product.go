package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name               string    `json:"name"`
	Product            []*Product
	SubCategories      []*ProductCategory `gorm:"foreignkey:ParentID"`
	ParentID           uuid.UUID
	BusinessUniverseID uuid.UUID
}

type ProductPicture struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	ProductID uuid.UUID
}

type Product struct {
	gorm.Model
	ID                 uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name               string    `json:"name"`
	ShortDescription   string    `json:"shortDescription"`
	Description        string    `json:"description"`
	Price              float64   `json:"price"`
	PriceIn            float64   `json:"priceIn"`
	Pictures           []*ProductPicture
	ProductCategoryID  uuid.UUID
	BusinessUniverseID uuid.UUID
}
