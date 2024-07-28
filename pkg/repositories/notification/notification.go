package notificationRepository

import (
	models "auth/auth_back/models/notification"
	"auth/auth_back/pkg/globalvars"
	"auth/auth_back/pkg/logger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var l = logger.Logger{}

type CreateMailConfrimationLink struct {
	UserId    uuid.UUID `json:"userId"`
	HashKey   string    `json:"hashKey"`
	SendTry   int       `json:"sendTry"`
	Confirmed bool      `json:"confirmed"`
}

type UpdateMailConfrimationLink struct {
	SendTry   int  `json:"sendTry"`
	Confirmed bool `json:"confirmed"`
}

type NotificationRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *NotificationRepository {
	notificationRepository := &NotificationRepository{
		db,
	}

	return notificationRepository
}

func (r NotificationRepository) CreateMailConfirmationLink(link *CreateMailConfrimationLink) (*models.MailConfrimationLinks, error) {

	newLink := &models.MailConfrimationLinks{
		UserId:    link.UserId,
		HashKey:   link.HashKey,
		SendTry:   link.SendTry,
		Confirmed: link.Confirmed,
	}

	tx := r.db.Create(&newLink)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/user/pg/repository.CreateItem")
		return nil, tx.Error
	}

	return newLink, nil
}

func (r NotificationRepository) FindLinkById(id uuid.UUID) *models.MailConfrimationLinks {
	findedLink := &models.MailConfrimationLinks{}

	tx := r.db.First(&findedLink, id)

	if tx.Error != nil {
		return nil
	}

	return findedLink
}

func (r NotificationRepository) FindLinkByUserId(id uuid.UUID) *models.MailConfrimationLinks {
	findedLink := &models.MailConfrimationLinks{}

	tx := r.db.Where("user_id = ?", id).First(&findedLink)

	if tx.Error != nil {
		return nil
	}

	return findedLink
}

func (r NotificationRepository) UpdateLink(link *UpdateMailConfrimationLink, id uuid.UUID) (*models.MailConfrimationLinks, error) {
	findedLink := r.FindLinkById(id)

	if findedLink == nil {
		l.LogNotify("Link with ID: "+id.String()+" not exist, failed update", "pkg/user/pg/repository.UpdateLink")
		return nil, globalvars.GetNotFoundErrors("MailConfirmationLink").Error
	}

	findedLink.SendTry = link.SendTry
	findedLink.Confirmed = link.Confirmed

	tx := r.db.Save(&findedLink)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("MailConfirmationLink").Error
	}

	return findedLink, nil
}
