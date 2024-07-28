package user

import (
	modelsUser "auth/auth_back/models/user"
)

func toUserResponse(in *modelsUser.User, role *modelsUser.Role) *User {
	return &User{
		Id:         in.ID.String(),
		FirstName:  in.FirstName,
		LastName:   in.LastName,
		SecondName: in.SecondName,
		Phone:      in.Phone,
		Email:      in.Email,
		Avatar:     in.Avatar,
		Role: &Role{
			Id:   role.ID.String(),
			Name: role.Name,
		},
	}
}
