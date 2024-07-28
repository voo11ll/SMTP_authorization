package businessUniverseRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/logger"
	"auth/auth_back/pkg/utils"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BusinessUniverse struct {
	Name   string `json:"name"`
	Domain string `bson:"domain"`
}

var l = logger.Logger{}

type BusinessUniverseRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *BusinessUniverseRepository {
	businessUniverseRepository := &BusinessUniverseRepository{
		db,
	}

	return businessUniverseRepository
}

func (r BusinessUniverseRepository) CreateItem(ctx context.Context) (*models.BusinessUniverse, string, error) {
	name := utils.GenerateStringDomainName(24)

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	newItem := &models.BusinessUniverse{
		Name:   name,
		Domain: name + viper.GetString("domain.templateUrl"),
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/businessUniverse/pg/repository.CreateItem")
		return nil, tx.Error.Error(), tx.Error
	}

	return newItem, "", nil
}

func (r BusinessUniverseRepository) FindItemByName(ctx context.Context, name string) *models.BusinessUniverse {
	findedItem := &models.BusinessUniverse{}

	tx := r.db.Where("name = ?", name).First(&findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r BusinessUniverseRepository) FindBusinessUniverseById(ctx context.Context, id uuid.UUID) *models.BusinessUniverse {
	findedItem := &models.BusinessUniverse{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r BusinessUniverseRepository) UpdateItem(ctx context.Context, item *BusinessUniverse, itemId uuid.UUID) (*models.BusinessUniverse, string, error) {
	findedItem := r.FindBusinessUniverseById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Business Universe with ID: "+itemId.String()+" not exist, failed update", "pkg/role/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("BusinessUniverse").Enum + ": " + "Business Universe with ID: " + itemId.String() + " not exist, failed update", globalvars.GetNotFoundErrors("Business Universe").Error
	}

	findedItem.Name = item.Name
	findedItem.Domain = item.Domain

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("BusinessUniverse").Enum + ": " + tx.Error.Error(), globalvars.GetUpdateErrors("Business Universe").Error
	}

	return findedItem, "Updated", nil
}
