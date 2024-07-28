package roleHandler

import (
	grpcService "auth/auth_back/pkg/services/grpc/user"
)

type CreateRequest struct {
	Name string `json:"name"`
}

func toCreateRequest(r *CreateRequest) *grpcService.RoleCreateRequest {
	return &grpcService.RoleCreateRequest{
		Name: r.Name,
	}
}

type GetRequest struct {
	Id string `json:"id"`
}

func toGetRequest(r *GetRequest) *grpcService.RoleGetRequest {
	return &grpcService.RoleGetRequest{
		Id: r.Id,
	}
}

type GetByNameRequest struct {
	Name string `json:"name"`
}

func toGetByNameRequest(r *GetByNameRequest) *grpcService.RoleGetByNameRequest {
	return &grpcService.RoleGetByNameRequest{
		Name: r.Name,
	}
}

type GetsRequest struct {
}

type UpdateRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func toUpdateRequest(r *UpdateRequest) *grpcService.RoleUpdateRequest {
	return &grpcService.RoleUpdateRequest{
		Id:   r.Id,
		Name: r.Name,
	}
}
