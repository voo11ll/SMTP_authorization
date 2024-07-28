package companyHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
	contactTypeHandler "auth/auth_back/pkg/services/http/organization/contactType"
)

type CompanyType struct {
	Id                 string             `json:"id"`
	Name               string             `json:"name"`
	FullName           string             `json:"fullName"`
	INN                int32              `json:"inn"`
	KPP                int32              `json:"kpp"`
	LegalAddress       string             `json:"legalAddress"`
	Banks              []*BankType        `json:"banks"`
	ContactInfos       []*ContactInfoType `json:"contactInfos"`
	BusinessUniverseId string             `json:"businessUniverseId"`
}

type CompanyResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Company CompanyType `json:"company"`
}

type CompaniesResponse struct {
	Code      int32          `json:"code"`
	Message   string         `json:"message"`
	Companies []*CompanyType `json:"companies"`
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
	Id          string                         `json:"id"`
	ContactType contactTypeHandler.ContactType `json:"contactType"`
	Value       string                         `json:"value"`
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

func ToCompanyResponse(r *org.CompanyResponse) *CompanyResponse {
	return &CompanyResponse{
		Code:    r.Code,
		Message: r.Message,
		Company: CompanyType{
			Id:                 r.Company.Id,
			Name:               r.Company.Name,
			FullName:           r.Company.FullName,
			INN:                r.Company.Inn,
			KPP:                r.Company.Kpp,
			LegalAddress:       r.Company.LegalAddress,
			Banks:              toBanksResponse(r.Company.Banks),
			ContactInfos:       toContactInfosResponse(r.Company.ContactInfos),
			BusinessUniverseId: r.Company.BusinessUniverse.Id,
		},
	}
}

func ToCompaniesResponse(r *org.CompaniesGetResponse) *CompaniesResponse {
	var out = make([]*CompanyType, len(r.Companies))

	for i, _u := range r.Companies {
		out[i] = &CompanyType{
			Id:                 _u.Id,
			Name:               _u.Name,
			FullName:           _u.FullName,
			INN:                _u.Inn,
			KPP:                _u.Kpp,
			LegalAddress:       _u.LegalAddress,
			Banks:              toBanksResponse(_u.Banks),
			ContactInfos:       toContactInfosResponse(_u.ContactInfos),
			BusinessUniverseId: _u.BusinessUniverse.Id,
		}
	}

	return &CompaniesResponse{
		Code:      r.Code,
		Message:   r.Message,
		Companies: out,
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
		Id:          r.Id,
		ContactType: toContactTypeResponse(r.ContactType),
		Value:       r.Value,
	}
}

func toContactTypeResponse(r *org.ContactType) contactTypeHandler.ContactType {
	return contactTypeHandler.ContactType{
		Id:     r.Id,
		Name:   r.Name,
		Notion: r.Notion,
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
