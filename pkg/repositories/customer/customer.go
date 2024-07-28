package customerRepository

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

type CustomerRepository struct {
	db *gorm.DB
}

func ConnectRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: dbs.InitDB(),
	}
}

func (r CustomerRepository) CreateItem(ctx context.Context, customer *Customer) (*models.Customer, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	businessUniverseId, _ := uuid.Parse(customer.BusinessUniverseId)

	newItem := &models.Customer{
		Name:               customer.Name,
		FullName:           customer.FullName,
		INN:                customer.INN,
		KPP:                customer.KPP,
		LegalAddress:       customer.LegalAddress,
		BusinessUniverseID: businessUniverseId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	for _, bank := range customer.Banks {
		bank.CustomerId = newItem.ID.String()
		_, err, _ := r.CreateBankItem(ctx, bank)
		if err != nil {
			l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
			return nil, tx.Error, tx.Error.Error()
		}
	}

	for _, contactInfo := range customer.ContactInfos {
		contactInfo.CustomerId = newItem.ID.String()
		_, err, _ := r.CreateContactInfoItem(ctx, contactInfo)

		if err != nil {
			l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
			return nil, tx.Error, tx.Error.Error()
		}
	}

	findedItem := r.FindItemById(ctx, newItem.ID)

	return findedItem, nil, "Created"
}

func (r CustomerRepository) FindItemById(ctx context.Context, id uuid.UUID) *models.Customer {
	findedItem := &models.Customer{}
	var findedBanks []*models.CustomerBank
	var findedContactInfos []*models.CustomerContactInfo

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

func (r CustomerRepository) FindAllItems(ctx context.Context, businessUniverseId uuid.UUID) []*models.Customer {
	var findedItems []*models.Customer

	l.LogNotify("Find Customers", "pkg/Customer/pg/repository.UpdateItem")

	tx := r.db.Where("business_universe_id = ?", businessUniverseId).Find(&findedItems)

	if tx.Error != nil {
		l.LogError(tx.Error.Error(), "pkg/Customer/pg/repository.UpdateItem")
		return nil
	}

	var findedBanks []*models.CustomerBank
	var findedContactInfos []*models.CustomerContactInfo

	var out = make([]*models.Customer, len(findedItems))

	for i, findedItem := range findedItems {
		findedBanks = r.FindAllBankItems(ctx, findedItem.ID)
		findedContactInfos = r.FindAllContactInfoItems(ctx, findedItem.ID)
		findedItem.Banks = findedBanks
		findedItem.ContactInfos = findedContactInfos
		out[i] = findedItem
	}

	return out
}

func (r CustomerRepository) UpdateItem(ctx context.Context, item *Customer, itemId uuid.UUID) (*models.Customer, error, string) {
	findedItem := r.FindItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Customer with ID: "+itemId.String()+" not exist, failed update", "pkg/Customer/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Customer").Error, globalvars.GetNotFoundErrors("Customer").Enum
	}

	findedItem.Name = item.Name
	findedItem.FullName = item.FullName
	findedItem.INN = item.INN
	findedItem.KPP = item.KPP
	findedItem.LegalAddress = item.LegalAddress

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer").Error, globalvars.GetUpdateErrors("Customer").Enum
	}

	return findedItem, nil, "Updated"
}
