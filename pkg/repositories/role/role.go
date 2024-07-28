package roleRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/user"
	dbs "auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/logger"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	Name string `json:"name"`
}

var l = logger.Logger{}

type RoleRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *RoleRepository {
	roleRepository := &RoleRepository{
		db: dbs.InitDB(),
	}

	roleExists := roleRepository.FindItemByName(context.TODO(), "Administrator")

	if roleExists == nil {
		roleRepository.CreateItem(context.TODO(), &Role{
			Name: "Administrator",
		})
	}

	return roleRepository
}

func (r RoleRepository) CreateItem(ctx context.Context, role *Role) (*models.Role, string, error) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	newItem := &models.Role{
		Name: role.Name,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/role/pg/repository.CreateItem")
		return nil, tx.Error.Error(), tx.Error
	}

	return newItem, "Created", nil
}

func (r RoleRepository) FindItemByName(ctx context.Context, name string) *models.Role {
	findedItem := &models.Role{}

	tx := r.db.Where("name = ?", name).First(&findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r RoleRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.Role {
	findedItem := &models.Role{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r RoleRepository) FindAllItems(ctx context.Context) []*models.Role {
	var findedItems []*models.Role

	tx := r.db.Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r RoleRepository) UpdateItem(ctx context.Context, item *Role, itemId uuid.UUID) (*models.Role, string, error) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Role with ID: "+itemId.String()+" not exist, failed update", "pkg/role/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Role").Enum, globalvars.GetNotFoundErrors("Role").Error
	}

	findedItem.Name = item.Name

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Role").Enum, globalvars.GetUpdateErrors("Role").Error
	}

	return findedItem, "Updated", nil
}
