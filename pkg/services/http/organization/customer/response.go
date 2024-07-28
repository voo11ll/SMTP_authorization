package customerHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type BusinessUniverseType struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

type CustomerType struct {
	Id               string               `json:"id"`
	Name             string               `json:"name"`
	FullName         string               `json:"fullName"`
	INN              int32                `json:"inn"`
	KPP              int32                `json:"kpp"`
	LegalAddress     string               `json:"legalAddress"`
	Banks            []*BankType          `json:"banks"`
	ContactInfos     []*ContactInfoType   `json:"contactInfos"`
	BusinessUniverse BusinessUniverseType `json:"businessUniverse"`
}

type CustomerResponse struct {
	Code     int32        `json:"code"`
	Message  string       `json:"message"`
	Customer CustomerType `json:"customer"`
}

type CustomersResponse struct {
	Code      int32           `json:"code"`
	Message   string          `json:"message"`
	Customers []*CustomerType `json:"customers"`
}

type BankType struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	AccountNumber     string `json:"accountNumber"`
	Bik               string `json:"bik"`
	CorrAccountNumber string `json:"corrAccountNumber"`
}

type GetBankResponse struct {
	Code    int32    `json:"code"`
	Message string   `json:"message"`
	Bank    BankType `json:"bank"`
}

type GetBanksResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Banks   []*BankType `json:"banks"`
}

type ContactInfoType struct {
	Id          string      `json:"id"`
	ContactType ContactType `json:"contactType"`
	Value       string      `json:"value"`
}

type ContactType struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Notion string `json:"notion"`
}

type GetContactInfoResponse struct {
	Code        int32           `json:"code"`
	Message     string          `json:"message"`
	ContactInfo ContactInfoType `json:"contactInfo"`
}

type GetContactInfosResponse struct {
	Code         int32              `json:"code"`
	Message      string             `json:"message"`
	ContactInfos []*ContactInfoType `json:"contactInfos"`
}

func ToCustomerResponse(r *org.CustomerResponse) *CustomerResponse {
	return &CustomerResponse{
		Code:    r.Code,
		Message: r.Message,
		Customer: CustomerType{
			Id:           r.Customer.Id,
			Name:         r.Customer.Name,
			FullName:     r.Customer.FullName,
			INN:          r.Customer.Inn,
			KPP:          r.Customer.Kpp,
			LegalAddress: r.Customer.LegalAddress,
			Banks:        toBanksResponse(r.Customer.Banks),
			ContactInfos: toContactInfosResponse(r.Customer.ContactInfos),
			BusinessUniverse: BusinessUniverseType{
				Id:     r.Customer.BusinessUniverse.Id,
				Name:   r.Customer.BusinessUniverse.Name,
				Domain: r.Customer.BusinessUniverse.Domain,
			},
		},
	}
}

func ToCustomersResponse(r *org.CustomersGetResponse) *CustomersResponse {
	var out = make([]*CustomerType, len(r.Customers))

	for i, _u := range r.Customers {
		out[i] = &CustomerType{
			Id:           _u.Id,
			Name:         _u.Name,
			FullName:     _u.FullName,
			INN:          _u.Inn,
			KPP:          _u.Kpp,
			LegalAddress: _u.LegalAddress,
			Banks:        toBanksResponse(_u.Banks),
			ContactInfos: toContactInfosResponse(_u.ContactInfos),
			BusinessUniverse: BusinessUniverseType{
				Id:     _u.BusinessUniverse.Id,
				Name:   _u.BusinessUniverse.Name,
				Domain: _u.BusinessUniverse.Domain,
			},
		}
	}

	return &CustomersResponse{
		Code:      r.Code,
		Message:   r.Message,
		Customers: out,
	}
}

func toBankResponse(r *org.Bank) *BankType {
	return &BankType{
		Id:                r.Id,
		Name:              r.Name,
		AccountNumber:     r.AccountNumber,
		Bik:               r.Bik,
		CorrAccountNumber: r.CorrAccountNumber,
	}
}

func ToGetBankResponse(r *org.GetBankResponse) *GetBankResponse {
	return &GetBankResponse{
		Code:    r.Code,
		Message: r.Message,
		Bank:    *toBankResponse(r.Bank),
	}
}

func toBanksResponse(r []*org.Bank) []*BankType {
	var out = make([]*BankType, len(r))

	for i, _u := range r {
		out[i] = toBankResponse(_u)
	}

	return out
}

func ToGetBanksResponse(r *org.GetBanksResponse) *GetBanksResponse {
	return &GetBanksResponse{
		Code:    r.Code,
		Message: r.Message,
		Banks:   toBanksResponse(r.Banks),
	}
}

func toContactInfoResponse(r *org.ContactInfo) *ContactInfoType {
	return &ContactInfoType{
		Id: r.Id,
		ContactType: ContactType{
			Id:     r.ContactType.Id,
			Name:   r.ContactType.Name,
			Notion: r.ContactType.Notion,
		},
		Value: r.Value,
	}
}

func toContactInfosResponse(r []*org.ContactInfo) []*ContactInfoType {
	var out = make([]*ContactInfoType, len(r))

	for i, _u := range r {
		out[i] = toContactInfoResponse(_u)
	}

	return out
}

func ToGetContactInfoResponse(r *org.GetContactInfoResponse) *GetContactInfoResponse {
	return &GetContactInfoResponse{
		Code:        r.Code,
		Message:     r.Message,
		ContactInfo: *toContactInfoResponse(r.ContactInfo),
	}
}

func ToGetContactInfosResponse(r *org.GetContactInfosResponse) *GetContactInfosResponse {
	return &GetContactInfosResponse{
		Code:         r.Code,
		Message:      r.Message,
		ContactInfos: toContactInfosResponse(r.ContactInfos),
	}
}
