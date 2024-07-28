package authHandler

import (
	"auth/auth_back/pkg/services/grpc/user"
	roleHandler "auth/auth_back/pkg/services/http/user/role"
)

type SignInResponse struct {
	Token   string   `json:"token"`
	Message string   `json:"message"`
	User    UserType `json:"user"`
}

func toSignInResponse(r *user.UserResponse, token string) *SignInResponse {
	return &SignInResponse{
		Token:   token,
		Message: r.Message,
		User: UserType{
			Id:         r.User.Id,
			FirstName:  r.User.FirstName,
			SecondName: r.User.SecondName,
			LastName:   r.User.LastName,
			Email:      r.User.Email,
			Phone:      r.User.Phone,
			Avatar:     r.User.Avatar,
			Role:       toRoleType(r.User.Role),
		},
	}
}

type UserType struct {
	Id         string               `json:"id"`
	FirstName  string               `json:"firstName"`
	LastName   string               `json:"lastName"`
	SecondName string               `json:"secondName"`
	Email      string               `json:"email"`
	Phone      string               `json:"phone"`
	Avatar     string               `json:"avatar"`
	Role       roleHandler.RoleType `json:"role"`
}

type EmailConfirmResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type SignUpResponse struct {
	Message string   `json:"message"`
	User    UserType `json:"user"`
}

func toSignUpResponse(r *user.UserResponse) *SignUpResponse {

	if r.User == nil {
		return &SignUpResponse{
			Message: r.Message,
			User:    UserType{},
		}
	}

	return &SignUpResponse{
		Message: r.Message,
		User: UserType{
			Id:         r.User.Id,
			FirstName:  r.User.FirstName,
			SecondName: r.User.SecondName,
			LastName:   r.User.LastName,
			Email:      r.User.Email,
			Phone:      r.User.Phone,
			Avatar:     r.User.Avatar,
			Role:       toRoleType(r.User.Role),
		},
	}
}

func toRoleType(r *user.Role) roleHandler.RoleType {
	if r == nil {
		return roleHandler.RoleType{}
	}

	return roleHandler.RoleType{
		Id:   r.Id,
		Name: r.Name,
	}
}
