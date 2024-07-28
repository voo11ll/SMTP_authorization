package customerRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	"auth/auth_back/pkg/globalvars"
	"context"
	"log"

	"github.com/google/uuid"
)

func (r CustomerRepository) CreateBankItem(ctx context.Context, bank *Bank) (*models.CustomerBank, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	customerId, _ := uuid.Parse(bank.CustomerId)

	newItem := &models.CustomerBank{
		Name:              bank.Name,
		AccountNumber:     bank.AccountNumber,
		Bik:               bank.Bik,
		CorrAccountNumber: bank.CorrAccountNumber,
		IsOpen:            bank.IsOpen,
		CustomerID:        customerId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r CustomerRepository) FindBankItemById(ctx context.Context, id uuid.UUID) *models.CustomerBank {
	findedItem := &models.CustomerBank{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CustomerRepository) FindAllBankItems(ctx context.Context, customerId uuid.UUID) []*models.CustomerBank {
	var findedItems []*models.CustomerBank

	tx := r.db.Where("customer_id = ?", customerId).Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r CustomerRepository) UpdateBankItem(ctx context.Context, item *Bank, itemId uuid.UUID) (*models.CustomerBank, error, string) {
	findedItem := r.FindBankItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Company Bank with ID: "+itemId.String()+" not exist, failed update", "pkg/bank/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Customer Bank").Error, globalvars.GetNotFoundErrors("CustomerBank").Enum
	}

	findedItem.Name = item.Name
	findedItem.AccountNumber = item.AccountNumber
	findedItem.Bik = item.Bik
	findedItem.CorrAccountNumber = item.CorrAccountNumber

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer Bank").Error, globalvars.GetUpdateErrors("CustomerBank").Enum
	}

	return findedItem, nil, "Updated"
}

func (r CustomerRepository) CloseBankItem(ctx context.Context, isOpen bool, itemId uuid.UUID) (*models.CustomerBank, error, string) {
	findedItem := r.FindBankItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Customer Bank with ID: "+itemId.String()+" not exist, failed update", "pkg/bank/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Customer Bank").Error, globalvars.GetNotFoundErrors("CustomerBank").Enum
	}

	findedItem.IsOpen = isOpen

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Customer Bank").Error, globalvars.GetUpdateErrors("CustomerBank").Enum
	}

	return findedItem, nil, "Updated"
}
