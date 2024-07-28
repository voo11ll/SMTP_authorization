package customerUserRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	dbs "auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/logger"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var l = logger.Logger{}

type User struct {
	FirstName          string `json:"firstname"`
	LastName           string `json:"lastname"`
	SecondName         string `json:"secondname"`
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	Password           string `json:"password"`
	Avatar             string `json:"avatar"`
	RoleID             string `json:"roleId"`
	CustomerID         string `json:"customerId"`
	BusinessUniverseID string `json:"businessUniverseId"`
}

type CustomerUserRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *CustomerUserRepository {
	return &CustomerUserRepository{
		db: dbs.InitDB(),
	}
}

func (r CustomerUserRepository) CreateItem(ctx context.Context, user *User) (*models.CustomerUser, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	roleId, _ := uuid.Parse(user.RoleID)
	businessUniverseId, _ := uuid.Parse(user.BusinessUniverseID)
	customerId, _ := uuid.Parse(user.CustomerID)

	newItem := &models.CustomerUser{
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		SecondName:         user.SecondName,
		Phone:              user.Phone,
		Email:              user.Email,
		Password:           user.Password,
		Avatar:             user.Avatar,
		RoleID:             roleId,
		CustomerID:         customerId,
		BusinessUniverseID: businessUniverseId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/user/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r CustomerUserRepository) FindItemByEmail(ctx context.Context, email string) *models.CustomerUser {
	findedItem := &models.CustomerUser{}

	tx := r.db.Where(&models.CustomerUser{Email: email}).First(&findedItem)

	fmt.Println(findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CustomerUserRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.CustomerUser {
	findedItem := &models.CustomerUser{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CustomerUserRepository) FindAllItems(ctx context.Context, customerId uuid.UUID) []*models.CustomerUser {
	var findedItems []*models.CustomerUser

	tx := r.db.Where("customer_id = ?", customerId).Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r CustomerUserRepository) UpdateItem(ctx context.Context, item *User, itemId uuid.UUID) (*models.CustomerUser, error, string) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Customer User with ID: "+itemId.String()+" not exist, failed update", "pkg/user/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Customer User").Error, globalvars.GetNotFoundErrors("CustomerUser").Enum
	}

	findedItem.FirstName = item.FirstName
	findedItem.LastName = item.LastName
	findedItem.SecondName = item.SecondName
	findedItem.Phone = item.Phone

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer User").Error, globalvars.GetUpdateErrors("CustomerUser").Enum
	}

	return findedItem, nil, "Updated"
}

func (r CustomerUserRepository) ChangeUserPassword(ctx context.Context, id uuid.UUID, password string) (*models.CustomerUser, error, string) {
	findedItem := &models.CustomerUser{}

	tx := r.db.First(&findedItem, id).Update("password", password)

	if tx.Error != nil {
		return nil, tx.Error, tx.Error.Error()
	}

	return findedItem, nil, "Password successfully changed"
}
