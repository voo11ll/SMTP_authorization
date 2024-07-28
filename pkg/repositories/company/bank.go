package companyRepository

import (
	"auth/auth_back/config"
	models "auth/auth_back/models/organization"
	"auth/auth_back/pkg/globalvars"
	"context"
	"log"

	"github.com/google/uuid"
)

func (r CompanyRepository) CreateBankItem(ctx context.Context, bank *Bank) (*models.CompanyBank, error, string) {

	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	companyId, _ := uuid.Parse(bank.CompanyId)

	newItem := &models.CompanyBank{
		Name:              bank.Name,
		AccountNumber:     bank.AccountNumber,
		Bik:               bank.Bik,
		CorrAccountNumber: bank.CorrAccountNumber,
		IsOpen:            bank.IsOpen,
		CompanyID:         companyId,
	}

	tx := r.db.Create(&newItem)

	if tx.Error != nil {
		l.LogNotify(tx.Error.Error(), "pkg/bank/pg/repository.CreateItem")
		return nil, tx.Error, tx.Error.Error()
	}

	return newItem, nil, "Created"
}

func (r CompanyRepository) FindBankItemById(ctx context.Context, id uuid.UUID) *models.CompanyBank {
	findedItem := &models.CompanyBank{}

	tx := r.db.First(&findedItem, id)

	if tx.Error != nil {
		return nil
	}

	return findedItem
}

func (r CompanyRepository) FindAllBankItems(ctx context.Context, companyId uuid.UUID) []*models.CompanyBank {
	var findedItems []*models.CompanyBank

	tx := r.db.Where("company_id = ?", companyId).Find(&findedItems)

	if tx.Error != nil {
		return nil
	}

	return findedItems
}

func (r CompanyRepository) UpdateBankItem(ctx context.Context, item *Bank, itemId uuid.UUID) (*models.CompanyBank, error, string) {
	findedItem := r.FindBankItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Company Bank with ID: "+itemId.String()+" not exist, failed update", "pkg/bank/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Company Bank").Error, globalvars.GetNotFoundErrors("CompanyBank").Enum
	}

	findedItem.Name = item.Name
	findedItem.AccountNumber = item.AccountNumber
	findedItem.Bik = item.Bik
	findedItem.CorrAccountNumber = item.CorrAccountNumber

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Company Bank").Error, globalvars.GetUpdateErrors("CompanyBank").Enum
	}

	return findedItem, nil, "Updated"
}

func (r CompanyRepository) CloseBankItem(ctx context.Context, isOpen bool, itemId uuid.UUID) (*models.CompanyBank, error, string) {
	findedItem := r.FindBankItemById(ctx, itemId)

	if findedItem == nil {
		l.LogNotify("Company Bank with ID: "+itemId.String()+" not exist, failed update", "pkg/bank/pg/repository.UpdateItem")
		return nil, globalvars.GetNotFoundErrors("Company Bank").Error, globalvars.GetNotFoundErrors("CompanyBank").Enum
	}

	findedItem.IsOpen = isOpen

	tx := r.db.Save(&findedItem)

	if tx.Error != nil {
		return nil, globalvars.GetUpdateErrors("Company Bank").Error, globalvars.GetUpdateErrors("CompanyBank").Enum
	}

	return findedItem, nil, "Updated"
}
