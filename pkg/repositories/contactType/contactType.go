package contactTypeRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	dbs "auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/logger"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactType struct {
	Name   string `json:"name"`
	Notion string `json:"notion"`
}

var l = logger.Logger{}

type ContactTypeRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *ContactTypeRepository {
	return &ContactTypeRepository{
		db: dbs.InitDB(),
	}
}

func (r ContactTypeRepository) CreateItem(ctx context.Context, contactType *ContactType) (*models.ContactType, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	newItem := &models.ContactType{
		Name:   contactType.Name,
		Notion: contactType.Notion,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/contactType/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r ContactTypeRepository) FindItemByName(ctx context.Context, name string) *models.ContactType {
	findedItem := &models.ContactType{}

	tx := r.db.Where("name = ?", name).First(&findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r ContactTypeRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.ContactType {
	findedItem := &models.ContactType{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r ContactTypeRepository) FindAllItems(ctx context.Context) []*models.ContactType {
	var findedItems []*models.ContactType

	tx := r.db.Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r ContactTypeRepository) UpdateItem(ctx context.Context, item *ContactType, itemId uuid.UUID) (*models.ContactType, error, string) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Contact type with ID: "+itemId.String()+" not exist, failed update", "pkg/contactType/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Contact Type").Error, globalvars.GetNotFoundErrors("ContactType").Enum
	}

	findedItem.Name = item.Name
	findedItem.Notion = item.Notion

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Contact Type").Error, globalvars.GetUpdateErrors("ContactType").Enum
	}

	return findedItem, nil, "Updated"
}
