package companyRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	"auth/auth_back/pkg/globalvars"
	"context"
	"log"

	"github.com/google/uuid"
)

func (r CompanyRepository) CreateContactInfoItem(ctx context.Context, contactInfo *ContactInfo) (*models.CompanyContactInfo, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	companyId, _ := uuid.Parse(contactInfo.CompanyId)
	contactTypeId, _ := uuid.Parse(contactInfo.ContactTypeID)

	newItem := &models.CompanyContactInfo{
		ContactTypeID: contactTypeId,
		Value:         contactInfo.Value,
		CompanyID:     companyId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/contactInfo/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r CompanyRepository) FindContactInfoItemById(ctx context.Context, id uuid.UUID) *models.CompanyContactInfo {
	findedItem := &models.CompanyContactInfo{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CompanyRepository) FindAllContactInfoItems(ctx context.Context, companyId uuid.UUID) []*models.CompanyContactInfo {
	var findedItems []*models.CompanyContactInfo

	tx := r.db.Where("company_id = ?", companyId).Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r CompanyRepository) UpdateContactInfoItem(ctx context.Context, item *ContactInfo, itemId uuid.UUID) (*models.CompanyContactInfo, error, string) {
	findedItem := r.FindContactInfoItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Contact Info with ID: "+itemId.String()+" not exist, failed update", "pkg/contactInfo/pg/repository.UpdateContactInfoItem")
		return nil, globalvars.GetNotFoundErrors("Contact Info").Error, globalvars.GetNotFoundErrors("ContactInfo").Enum
	}

	contactTypeId, _ := uuid.Parse(item.ContactTypeID)

	findedItem.ContactTypeID = contactTypeId
	findedItem.Value = item.Value

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Contact Info").Error, globalvars.GetUpdateErrors("ContactInfo").Enum
	}

	return findedItem, nil, "Updated"
}

func (r CompanyRepository) DeleteContactInfoItem(ctx context.Context, itemId uuid.UUID) (*models.CompanyContactInfo, error, string) {
	findedItem := r.FindContactInfoItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Contact Info with ID: "+itemId.String()+" not exist, failed deleting", "pkg/contactInfo/pg/repository.UpdateContactInfoItem")
		return nil, globalvars.GetNotFoundErrors("Contact Info").Error, globalvars.GetNotFoundErrors("ContactInfo").Enum
	}

	tx := r.db.Delete(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Contact Info").Error, globalvars.GetUpdateErrors("ContactInfo").Enum
	}

	return findedItem, nil, "Deleted"
}
