package customerUserHandler

import (
	"auth/auth_back/pkg/services/grpc/user"
)

type AddUserRequest struct {
	Email              string `json:"email" validate:"required"`
	Password           string `json:"password" validate:"required"`
	CustomerId         string `json:"customerId" validate:"required"`
	BusinessUniverseId string `json:"businessUniverseId" validate:"required"`
}

func toAddUserRequest(r *AddUserRequest, roleId string) *user.CustomerUserAddRequest {
	return &user.CustomerUserAddRequest{
		Email:              r.Email,
		Password:           r.Password,
		RoleID:             roleId,
		CustomerID:         r.CustomerId,
		BusinessUniverseID: r.BusinessUniverseId,
	}
}

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

func toUserGetByEmailRequest(r *GetUserByEmailRequest) *user.CustomerUserGetByEmailRequest {
	return &user.CustomerUserGetByEmailRequest{
		Email: r.Email,
	}
}

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

func toUserGetByIdRequest(r *GetUserByIdRequest) *user.CustomerUserGetByIdRequest {
	return &user.CustomerUserGetByIdRequest{
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

func toUserUpdateRequest(r *UpdateRequest) *user.CustomerUserUpdateRequest {
	return &user.CustomerUserUpdateRequest{
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

func toChangePasswordRequest(r *ChangePasswordRequest) *user.CustomerUserChangePasswordRequest {
	return &user.CustomerUserChangePasswordRequest{
		Id:       r.Id,
		Password: r.Password,
	}
}
