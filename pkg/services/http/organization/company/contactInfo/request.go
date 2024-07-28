package companyHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type CreateContactInfoRequest struct {
	ContactTypeId string `json:"contactTypeId"`
	Value         string `json:"value"`
	CompanyId     string `json:"companyId"`
}

type GetContactInfoRequest struct {
	Id string `json:"id"`
}

type GetContactInfosRequest struct {
	CompanyId string `json:"companyId"`
}

type UpdateContactInfoRequest struct {
	Id            string `json:"id"`
	ContactTypeId string `json:"contactTypeId"`
	Value         string `json:"value"`
}

func toGrpcCreateCompanyContactInfoRequest(in *CreateContactInfoRequest) *org.ContactInfoCreateType {
	return &org.ContactInfoCreateType{
		ContactTypeID: in.ContactTypeId,
		Value:         in.Value,
	}
}

func ToGrpcCreateCompanyContactInfosRequest(in []*CreateContactInfoRequest) []*org.ContactInfoCreateType {
	var out = make([]*org.ContactInfoCreateType, len(in))

	for i, _u := range in {
		out[i] = toGrpcCreateCompanyContactInfoRequest(_u)
	}

	return out
}

func toGetContactInfoRequest(r *GetContactInfoRequest) *org.GetContactInfoRequest {
	return &org.GetContactInfoRequest{
		Id: r.Id,
	}
}

func toGetContactInfosRequest(r *GetContactInfosRequest) *org.GetCompanyContactInfosRequest {
	return &org.GetCompanyContactInfosRequest{
		CompanyId: r.CompanyId,
	}
}
