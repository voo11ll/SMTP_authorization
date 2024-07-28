package user

import (
	models "auth/auth_back/models/user"
	roleRepository "auth/auth_back/pkg/repositories/role"
)

func toRoleCreateModel(in *RoleCreateRequest) *roleRepository.Role {
	return &roleRepository.Role{
		Name: in.GetName(),
	}
}

func toRolesGet(u []*models.Role) []*Role {
	var out = make([]*Role, len(u))

	for i, _u := range u {
		out[i] = &Role{
			Id:   _u.ID.String(),
			Name: _u.Name,
		}
	}

	return out
}

func toRoleUpdateModel(in *RoleUpdateRequest) *roleRepository.Role {
	return &roleRepository.Role{
		Name: in.GetName(),
	}
}
