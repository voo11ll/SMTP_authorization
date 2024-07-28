package companyRepository

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

var l = logger.Logger{}

type CompanyRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		db: dbs.InitDB(),
	}
}

func (r CompanyRepository) CreateItem(ctx context.Context, company *Company) (*models.Company, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	businessUniverseId, _ := uuid.Parse(company.BusinessUniverseId)

	newItem := &models.Company{
		Name:               company.Name,
		FullName:           company.FullName,
		INN:                company.INN,
		KPP:                company.KPP,
		LegalAddress:       company.LegalAddress,
		BusinessUniverseID: businessUniverseId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	for _, bank := range company.Banks {
		bank.CompanyId = newItem.ID.String()
		_, err, _ := r.CreateBankItem(ctx, bank)
		if err != nil {
			l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
			return nil, tx.Error, tx.Error.Error()
		}
	}

	for _, contactInfo := range company.ContactInfos {
		contactInfo.CompanyId = newItem.ID.String()
		_, err, _ := r.CreateContactInfoItem(ctx, contactInfo)

		if err != nil {
			l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
			return nil, tx.Error, tx.Error.Error()
		}
	}

	findedItem := r.FindItemById(ctx, newItem.ID)

	return findedItem, nil, "Created"
}

func (r CompanyRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.Company {
	findedItem := &models.Company{}
	var findedBanks []*models.CompanyBank
	var findedContactInfos []*models.CompanyContactInfo

	tx := r.db.First(&findedItem, id)
	findedBanks = r.FindAllBankItems(ctx, findedItem.ID)
	findedContactInfos = r.FindAllContactInfoItems(ctx, findedItem.ID)
	findedItem.Banks = findedBanks
	findedItem.ContactInfos = findedContactInfos

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CompanyRepository) FindAllItems(ctx context.Context, businessUniverseId uuid.UUID) []*models.Company {
	var findedItems []*models.Company

	l.LogNotify("Find companies", "pkg/Company/pg/repository.UpdateItem")

	tx := r.db.Where("business_universe_id = ?", businessUniverseId).Find(&findedItems)

	if tx.Error != nil {
		l.LogError(tx.Error.Error(), "pkg/Company/pg/repository.UpdateItem")
		return nil
	}

	var findedBanks []*models.CompanyBank
	var findedContactInfos []*models.CompanyContactInfo

	var out = make([]*models.Company, len(findedItems))

	for i, findedItem := range findedItems {
		findedBanks = r.FindAllBankItems(ctx, findedItem.ID)
		findedContactInfos = r.FindAllContactInfoItems(ctx, findedItem.ID)
		findedItem.Banks = findedBanks
		findedItem.ContactInfos = findedContactInfos
		out[i] = findedItem
	}

	return out
}

func (r CompanyRepository) UpdateItem(ctx context.Context, item *Company, itemId uuid.UUID) (*models.Company, error, string) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Company with ID: "+itemId.String()+" not exist, failed update", "pkg/Company/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Company").Error, globalvars.GetNotFoundErrors("Company").Enum
	}

	findedItem.Name = item.Name
	findedItem.FullName = item.FullName
	findedItem.INN = item.INN
	findedItem.KPP = item.KPP
	findedItem.LegalAddress = item.LegalAddress

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Company").Error, globalvars.GetUpdateErrors("Company").Enum
	}

	return findedItem, nil, "Updated"
}
