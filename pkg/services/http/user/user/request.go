package userHandler

import (
	"auth/auth_back/pkg/services/grpc/user"
)

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

func toUserGetByEmailRequest(r *GetUserByEmailRequest) *user.UserGetByEmailRequest {
	return &user.UserGetByEmailRequest{
		Email: r.Email,
	}
}

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

func toUserGetByIdRequest(r *GetUserByIdRequest) *user.UserGetByIdRequest {
	return &user.UserGetByIdRequest{
		Id: r.Id,
	}
}

type UpdateRequest struct {
	Id         string `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SecondName string `json:"secondName"`
	Phone      string `json:"phone"`
}

func toUserUpdateRequest(r *UpdateRequest) *user.UserUpdateRequest {
	return &user.UserUpdateRequest{
		Id:         r.Id,
		FirstName:  r.FirstName,
		LastName:   r.LastName,
		SecondName: r.SecondName,
		Phone:      r.Phone,
	}
}

type ChangePasswordRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func toChangePasswordRequest(r *ChangePasswordRequest) *user.UserChangePasswordRequest {
	return &user.UserChangePasswordRequest{
		Id:       r.Id,
		Password: r.Password,
	}
}
