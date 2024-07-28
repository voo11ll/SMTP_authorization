package user

import (
	"auth/auth_back/pkg/globalvars"
	context "context"

	uuid "github.com/google/uuid"
)

func (s *GrpcServer) CreateRole(ctx context.Context, in *RoleCreateRequest) (*RoleResponse, error) {
	tmpl := toRoleCreateModel(in)

	item, message, err := s.RoleRepo.CreateItem(context.TODO(), tmpl)

	if err != nil && message == "" {
		return &RoleResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on creating role",
			Role:    nil,
		}, nil
	} else if err != nil && message != "" {
		return &RoleResponse{
			Code:    globalvars.NotFound,
			Message: message,
			Role:    nil,
		}, nil
	} else {
		return &RoleResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Role: &Role{
				Id:   item.ID.String(),
				Name: item.Name,
			},
		}, nil
	}
}

func (s *GrpcServer) UpdateRole(ctx context.Context, in *RoleUpdateRequest) (*RoleResponse, error) {
	tmpl := toRoleUpdateModel(in)
	id, _ := uuid.Parse(in.Id)

	item, message, err := s.RoleRepo.UpdateItem(context.TODO(), tmpl, id)

	if err != nil && message == "" {
		return &RoleResponse{
			Code:    globalvars.ServerInternalError,
			Message: "Internal server error on updating role",
			Role:    nil,
		}, nil
	} else if err != nil && message != "" {
		return &RoleResponse{
			Code:    globalvars.NotFound,
			Message: message,
			Role:    nil,
		}, nil
	} else {
		return &RoleResponse{
			Code:    globalvars.StatusOK,
			Message: message,
			Role: &Role{
				Id:   item.ID.String(),
				Name: item.Name,
			},
		}, nil
	}
}

func (s *GrpcServer) GetRole(ctx context.Context, in *RoleGetRequest) (*RoleResponse, error) {

	id, _ := uuid.Parse(in.Id)

	role := s.RoleRepo.FindItemById(ctx, id)

	return &RoleResponse{
		Role: &Role{
			Id:   in.Id,
			Name: role.Name,
		},
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetRoleByName(ctx context.Context, in *RoleGetByNameRequest) (*RoleResponse, error) {

	item := s.RoleRepo.FindItemByName(ctx, in.Name)

	return &RoleResponse{
		Role: &Role{
			Id:   item.ID.String(),
			Name: in.Name,
		},
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}

func (s *GrpcServer) GetRoles(ctx context.Context, in *RolesGetRequest) (*RolesGetResponse, error) {

	items := s.RoleRepo.FindAllItems(ctx)

	return &RolesGetResponse{
		Roles:   toRolesGet(items),
		Code:    globalvars.StatusOK,
		Message: "OK",
	}, nil
}
