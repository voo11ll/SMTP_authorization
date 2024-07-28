package contactTypeHandler

import (
	grpcService "auth/auth_back/pkg/services/grpc/organization"
)

type ContactTypeResponse struct {
	Code        int32       `json:"code"`
	Message     string      `json:"message"`
	ContactType ContactType `json:"contactType"`
}

func toContactTypeResponse(r *grpcService.ContactTypeResponse) *ContactTypeResponse {
	return &ContactTypeResponse{
		Code:        r.Code,
		Message:     r.Message,
		ContactType: toItemType(r.ContactType),
	}
}

type ContactType struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Notion string `json:"notion"`
}

type GetsResponse struct {
	Code         int32         `json:"code"`
	Message      string        `json:"message"`
	ContactTypes []ContactType `json:"contactTypes"`
}

func toItemType(r *grpcService.ContactType) ContactType {
	return ContactType{
		Id:     r.Id,
		Name:   r.Name,
		Notion: r.Notion,
	}
}

func toGetsResponse(r *grpcService.ContactTypesResponse) *GetsResponse {
	out := make([]ContactType, len(r.ContactTypes))

	for i, _r := range r.ContactTypes {
		out[i] = toItemType(_r)
	}

	return &GetsResponse{
		Message:      r.Message,
		Code:         r.Code,
		ContactTypes: out,
	}
}
