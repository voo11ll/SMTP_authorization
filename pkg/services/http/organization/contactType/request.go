package contactTypeHandler

import (
	grpcService "auth/auth_back/pkg/services/grpc/organization"
)

type CreateRequest struct {
	Name   string `json:"name"`
	Notion string `json:"notion"`
}

func toCreateRequest(r *CreateRequest) *grpcService.ContactTypeCreateRequest {
	return &grpcService.ContactTypeCreateRequest{
		Name:   r.Name,
		Notion: r.Notion,
	}
}

type GetRequest struct {
	Id string `json:"id"`
}

func toGetRequest(r *GetRequest) *grpcService.ContactTypeGetRequest {
	return &grpcService.ContactTypeGetRequest{
		Id: r.Id,
	}
}

type GetsRequest struct {
}

type UpdateRequest struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Notion string `json:"notion"`
}

func toUpdateRequest(r *UpdateRequest) *grpcService.ContactTypeUpdateRequest {
	return &grpcService.ContactTypeUpdateRequest{
		Id:     r.Id,
		Name:   r.Name,
		Notion: r.Notion,
	}
}
