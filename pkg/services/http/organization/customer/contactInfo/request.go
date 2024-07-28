package customerHandler

import (
	org "auth/auth_back/pkg/services/grpc/organization"
)

type CreateContactInfoRequest struct {
	ContactTypeId string `json:"contactTypeId"`
	Value         string `json:"value"`
	CustomerId    string `json:"сustomerId"`
}

type GetContactInfoRequest struct {
	Id string `json:"id"`
}

type GetContactInfosRequest struct {
	CustomerId string `json:"сustomerId"`
}

type UpdateContactInfoRequest struct {
	Id            string `json:"id"`
	ContactTypeId string `json:"contactTypeId"`
	Value         string `json:"value"`
}

func toGrpcCreateCustomerContactInfoRequest(in *CreateContactInfoRequest) *org.ContactInfoCreateType {
	return &org.ContactInfoCreateType{
		ContactTypeID: in.ContactTypeId,
		Value:         in.Value,
	}
}

func ToGrpcCreateCustomerContactInfosRequest(in []*CreateContactInfoRequest) []*org.ContactInfoCreateType {
	var out = make([]*org.ContactInfoCreateType, len(in))

	for i, _u := range in {
		out[i] = toGrpcCreateCustomerContactInfoRequest(_u)
	}

	return out
}

func toGetContactInfoRequest(r *GetContactInfoRequest) *org.GetContactInfoRequest {
	return &org.GetContactInfoRequest{
		Id: r.Id,
	}
}

func toGetContactInfosRequest(r *GetContactInfosRequest) *org.GetCustomerContactInfosRequest {
	return &org.GetCustomerContactInfosRequest{
		CustomerId: r.CustomerId,
	}
}
