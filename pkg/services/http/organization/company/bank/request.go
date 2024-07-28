package companyHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type CreateBankRequest struct {
	Name              string `json:"name"`
	AccountNumber     string `json:"accountNumber"`
	Bik               string `json:"bik"`
	CorrAccountNumber string `json:"corrAccountNumber"`
	CompanyId         string `json:"companyId"`
}

func toGrpcCreateCompanyBankRequest(in *CreateBankRequest) *org.Bank {
	return &org.Bank{
		Name:              in.Name,
		AccountNumber:     in.AccountNumber,
		Bik:               in.Bik,
		CorrAccountNumber: in.CorrAccountNumber,
	}
}

func ToGrpcCreateCompanyBanksRequest(in []*CreateBankRequest) []*org.Bank {
	var out = make([]*org.Bank, len(in))

	for i, _u := range in {
		out[i] = toGrpcCreateCompanyBankRequest(_u)
	}

	return out
}

type GetRequest struct {
	Id string `json:"id"`
}

type GetBanksRequest struct {
	CompanyId string `json:"companyId"`
}

type UpdateRequest struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	AccountNumber     string `json:"accountNumber"`
	Bik               string `json:"bik"`
	CorrAccountNumber string `json:"corrAccountNumber"`
	IsOpen            bool   `json:"isOpen"`
}

func toGetRequest(r *GetRequest) *org.GetBankRequest {
	return &org.GetBankRequest{
		Id: r.Id,
	}
}

func toGetBanksRequest(r *GetBanksRequest) *org.GetCompanyBanksRequest {
	return &org.GetCompanyBanksRequest{
		CompanyId: r.CompanyId,
	}
}
