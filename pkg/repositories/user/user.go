package userRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/user"
	dbs "auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/helpers/passwordHelper"
	"auth/auth_back/pkg/logger"
	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
	notificationRepository "auth/auth_back/pkg/repositories/notification"
	roleRepository "auth/auth_back/pkg/repositories/role"
	"context"
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
	BusinessUniverseID string `json:"businessUniverseId"`
}

type UserRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *UserRepository {
	userRepository := &UserRepository{
		db: dbs.InitDB(),
	}

	var ctx context.Context

	userExists := userRepository.FindItemByEmail(ctx, "user@test.ru")

	if userExists == nil {

		buRepo := businessUniverseRepository.ConnectRepository(db)
		roleRepo := roleRepository.ConnectRepository(db)
		notifyRepo := notificationRepository.ConnectRepository(db)

		bUniverce, _, _ := buRepo.CreateItem(ctx)
		role := roleRepo.FindItemByName(ctx, "Administrator")

		userPass, _ := passwordHelper.HashPassword("qweasd")

		_, _, err := userRepository.CreateItem(ctx, &User{
			Email:              "user@test.ru",
			FirstName:          "User",
			LastName:           "Test",
			SecondName:         "test",
			Phone:              "1234567",
			Password:           userPass,
			RoleID:             role.ID.String(),
			BusinessUniverseID: bUniverce.ID.URN(),
		})

		if err != nil {
			l.LogError("Test user has not been created with error: ", err.Error())
		}

		cretedUser := userRepository.FindItemByEmail(ctx, "user@test.ru")

		key, _ := passwordHelper.HashPassword("user@test.ru")

		_, err = notifyRepo.CreateMailConfirmationLink(&notificationRepository.CreateMailConfrimationLink{
			UserId:    cretedUser.ID,
			HashKey:   key,
			SendTry:   1,
			Confirmed: true,
		})

		if err != nil {
			l.LogError(err.Error(), "pkg/user/pg/repository.CreateItem")
		}
	}

	return userRepository
}

func (r UserRepository) CreateItem(ctx context.Context, user *User) (*models.User, string, error) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	roleId, _ := uuid.Parse(user.RoleID)

	newItem := &models.User{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		SecondName: user.SecondName,
		Phone:      user.Phone,
		Email:      user.Email,
		Password:   user.Password,
		// Avatar:             user.Avatar,
		RoleID: roleId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/user/pg/repository.CreateItem")
		return nil, tx.Error.Error(), tx.Error
	}

	return newItem, "", nil
}

func (r UserRepository) FindItemByEmail(ctx context.Context, email string) *models.User {
	findedItem := &models.User{}

	tx := r.db.Where(&models.User{Email: email}).First(&findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r UserRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.User {
	findedItem := &models.User{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

// func (r UserRepository) FindAllItems(ctx context.Context, businessUniverseId uuid.UUID) []*models.User {
// 	var findedItems []*models.User

// 	tx := r.db.Where("business_universe_id = ?", businessUniverseId).Find(&findedItems)

// 	if tx.Error != nil {
// 		return nil
// 	}

// 	return findedItems
// }

func (r UserRepository) UpdateItem(ctx context.Context, item *User, itemId uuid.UUID) (*models.User, string, error) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("User with ID: "+itemId.String()+" not exist, failed update", "pkg/user/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("User").Enum, globalvars.GetNotFoundErrors("User").Error
	}

	findedItem.FirstName = item.FirstName
	findedItem.LastName = item.LastName
	findedItem.SecondName = item.SecondName
	findedItem.Phone = item.Phone

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("User").Enum, globalvars.GetUpdateErrors("User").Error
	}

	return findedItem, "", nil
}

func (r UserRepository) ChangeUserPassword(ctx context.Context, id uuid.UUID, password string) (*models.User, string, error) {
	findedItem := &models.User{}

	tx := r.db.First(&findedItem, id).Update("password", password)

	if tx.Error != nil {
		return nil, tx.Error.Error(), tx.Error
	}

	return findedItem, "Password successfully changed", nil
}

func (r UserRepository) CheckUser(ctx context.Context, email, password string) *models.User {

	var findedItem *models.User

	tx := r.db.Where("email = ? AND password = ?", email, password).Find(&findedItem)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}
