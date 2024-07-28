package customerRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	"auth/auth_back/pkg/globalvars"
	"context"
	"log"

	"github.com/google/uuid"
)

func (r CustomerRepository) CreateContactInfoItem(ctx context.Context, contactInfo *ContactInfo) (*models.CustomerContactInfo, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	customerId, _ := uuid.Parse(contactInfo.CustomerId)
	contactTypeId, _ := uuid.Parse(contactInfo.ContactTypeID)

	newItem := &models.CustomerContactInfo{
		ContactTypeID: contactTypeId,
		Value:         contactInfo.Value,
		CustomerID:    customerId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/contactInfo/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r CustomerRepository) FindContactInfoItemById(ctx context.Context, id uuid.UUID) *models.CustomerContactInfo {
	findedItem := &models.CustomerContactInfo{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CustomerRepository) FindAllContactInfoItems(ctx context.Context, customerId uuid.UUID) []*models.CustomerContactInfo {
	var findedItems []*models.CustomerContactInfo

	tx := r.db.Where("customer_id = ?", customerId).Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r CustomerRepository) UpdateContactInfoItem(ctx context.Context, item *ContactInfo, itemId uuid.UUID) (*models.CustomerContactInfo, error, string) {
	findedItem := r.FindContactInfoItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Contact Info with ID: "+itemId.String()+" not exist, failed update", "pkg/contactInfo/pg/repository.UpdateContactInfoItem")
		return nil, globalvars.GetNotFoundErrors("CustomerContact Info").Error, globalvars.GetNotFoundErrors("CustomerContactInfo").Enum
	}

	contactTypeId, _ := uuid.Parse(item.ContactTypeID)

	findedItem.ContactTypeID = contactTypeId
	findedItem.Value = item.Value

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer Contact Info").Error, globalvars.GetUpdateErrors("CustomerContactInfo").Enum
	}

	return findedItem, nil, "Updated"
}

func (r CustomerRepository) DeleteContactInfoItem(ctx context.Context, itemId uuid.UUID) (*models.CustomerContactInfo, error, string) {
	findedItem := r.FindContactInfoItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Contact Info with ID: "+itemId.String()+" not exist, failed deleting", "pkg/contactInfo/pg/repository.UpdateContactInfoItem")
		return nil, globalvars.GetNotFoundErrors("Contact Info").Error, globalvars.GetNotFoundErrors("ContactInfo").Enum
	}

	tx := r.db.Delete(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer Contact Info").Error, globalvars.GetUpdateErrors("CustomerContactInfo").Enum
	}

	return findedItem, nil, "Deleted"
}
