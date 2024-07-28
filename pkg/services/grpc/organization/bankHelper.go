package organization

import (
	models "auth/auth_back/models/organization"
	companyRepository "auth/auth_back/pkg/repositories/company"
	customerRepository "auth/auth_back/pkg/repositories/customer"
)

/**
* Company bank
 */
func toBankCreateModel(in *AddCompanyBankRequest) *companyRepository.Bank {
	return &companyRepository.Bank{
		Name:              in.Bank.Name,
		AccountNumber:     in.Bank.AccountNumber,
		Bik:               in.Bank.Bik,
		CorrAccountNumber: in.Bank.CorrAccountNumber,
		IsOpen:            true,
		CompanyId:         in.CompanyId,
	}
}

func toBanksCreateModel(in []*Bank) []*companyRepository.Bank {
	var out = make([]*companyRepository.Bank, len(in))

	for i, _u := range in {
		out[i] = &companyRepository.Bank{
			Name:              _u.Name,
			AccountNumber:     _u.AccountNumber,
			Bik:               _u.Bik,
			CorrAccountNumber: _u.CorrAccountNumber,
			IsOpen:            _u.IsOpen,
		}
	}

	return out
}

func toBanksGet(u []*models.CompanyBank) []*Bank {
	var out = make([]*Bank, len(u))

	for i, _u := range u {
		out[i] = &Bank{
			Id:                _u.ID.String(),
			Name:              _u.Name,
			AccountNumber:     _u.AccountNumber,
			Bik:               _u.Bik,
			CorrAccountNumber: _u.CorrAccountNumber,
			IsOpen:            _u.IsOpen,
		}
	}

	return out
}

func toBankGet(u *models.CompanyBank) *Bank {
	return &Bank{
		Id:                u.ID.String(),
		Name:              u.Name,
		AccountNumber:     u.AccountNumber,
		Bik:               u.Bik,
		CorrAccountNumber: u.CorrAccountNumber,
		IsOpen:            u.IsOpen,
	}
}

func toBankUpdateModel(in *UpdateBankRequest) *companyRepository.Bank {
	return &companyRepository.Bank{
		Name:              in.GetName(),
		AccountNumber:     in.AccountNumber,
		Bik:               in.Bik,
		CorrAccountNumber: in.CorrAccountNumber,
		IsOpen:            true,
	}
}

/**
* Customer bank
**/

func toCustomerBankCreateModel(in *AddCustomerBankRequest) *customerRepository.Bank {
	return &customerRepository.Bank{
		Name:              in.Bank.Name,
		AccountNumber:     in.Bank.AccountNumber,
		Bik:               in.Bank.Bik,
		CorrAccountNumber: in.Bank.CorrAccountNumber,
		IsOpen:            true,
		CustomerId:        in.CustomerId,
	}
}

func toCustomerBanksCreateModel(in []*Bank) []*customerRepository.Bank {
	var out = make([]*customerRepository.Bank, len(in))

	for i, _u := range in {
		out[i] = &customerRepository.Bank{
			Name:              _u.Name,
			AccountNumber:     _u.AccountNumber,
			Bik:               _u.Bik,
			CorrAccountNumber: _u.CorrAccountNumber,
			IsOpen:            _u.IsOpen,
		}
	}

	return out
}

func toCustomerBanksGet(u []*models.CustomerBank) []*Bank {
	var out = make([]*Bank, len(u))

	for i, _u := range u {
		out[i] = &Bank{
			Id:                _u.ID.String(),
			Name:              _u.Name,
			AccountNumber:     _u.AccountNumber,
			Bik:               _u.Bik,
			CorrAccountNumber: _u.CorrAccountNumber,
			IsOpen:            _u.IsOpen,
		}
	}

	return out
}

func toCustomerBankGet(u *models.CustomerBank) *Bank {
	return &Bank{
		Id:                u.ID.String(),
		Name:              u.Name,
		AccountNumber:     u.AccountNumber,
		Bik:               u.Bik,
		CorrAccountNumber: u.CorrAccountNumber,
		IsOpen:            u.IsOpen,
	}
}

func toCustomerBankUpdateModel(in *UpdateBankRequest) *customerRepository.Bank {
	return &customerRepository.Bank{
		Name:              in.GetName(),
		AccountNumber:     in.AccountNumber,
		Bik:               in.Bik,
		CorrAccountNumber: in.CorrAccountNumber,
		IsOpen:            true,
	}
}
