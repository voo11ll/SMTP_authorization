package roleHandler

import (
	grpcService "auth/auth_back/pkg/services/grpc/user"
)

type RoleResponse struct {
	Code    int32    `json:"code"`
	Message string   `json:"message"`
	Role    RoleType `json:"role"`
}

func toRoleResponse(r *grpcService.RoleResponse) *RoleResponse {
	return &RoleResponse{
		Code:    r.Code,
		Message: r.Message,
		Role:    toItemType(r.Role),
	}
}

type RoleType struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetsResponse struct {
	Code    int32      `json:"code"`
	Message string     `json:"message"`
	Roles   []RoleType `json:"roles"`
}

func toItemType(r *grpcService.Role) RoleType {
	return RoleType{
		Id:   r.Id,
		Name: r.Name,
	}
}

func toGetsResponse(r *grpcService.RolesGetResponse) *GetsResponse {
	out := make([]RoleType, len(r.Roles))

	for i, _r := range r.Roles {
		out[i] = toItemType(_r)
	}

	return &GetsResponse{
		Message: r.Message,
		Code:    r.Code,
		Roles:   out,
	}
}
