package companyHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
	bankHandler "auth/auth_back/pkg/services/http/organization/company/bank"
	contactInfoHandler "auth/auth_back/pkg/services/http/organization/company/contactInfo"
)

type CreateCompanyRequest struct {
	Name               string                                         `json:"name"`
	FullName           string                                         `json:"fullName"`
	INN                int32                                          `json:"inn"`
	KPP                int32                                          `json:"kpp"`
	LegalAddress       string                                         `json:"legalAddress"`
	Banks              []*bankHandler.CreateBankRequest               `json:"banks"`
	ContactInfos       []*contactInfoHandler.CreateContactInfoRequest `json:"contactInfos"`
	BusinessUniverseId string                                         `json:"businessUniverseId"`
}

type GetRequest struct {
	Id string `json:"id"`
}

type GetCompaniesRequest struct {
	BusinessUniverseId string `json:"businessUniverseId"`
}

type UpdateRequest struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	FullName     string `json:"fullName"`
	INN          int32  `json:"inn"`
	KPP          int32  `json:"kpp"`
	LegalAddress string `json:"legalAddress"`
}

func toGrpcCreateCompanyRequest(in *CreateCompanyRequest) *org.CreateRequest {
	return &org.CreateRequest{
		Name:               in.Name,
		FullName:           in.FullName,
		Inn:                in.INN,
		Kpp:                in.KPP,
		LegalAddress:       in.LegalAddress,
		BusinessUniverseID: in.BusinessUniverseId,
		Banks:              bankHandler.ToGrpcCreateCompanyBanksRequest(in.Banks),
		ContactInfos:       contactInfoHandler.ToGrpcCreateCompanyContactInfosRequest(in.ContactInfos),
	}
}

func toGetRequest(r *GetRequest) *org.GetRequest {
	return &org.GetRequest{
		Id: r.Id,
	}
}

func toGetCompaniesRequest(r *GetCompaniesRequest) *org.GetsRequest {
	return &org.GetsRequest{
		BusinessUniverseId: r.BusinessUniverseId,
	}
}
