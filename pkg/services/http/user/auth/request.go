package authHandler

import (
	"auth/auth_back/pkg/services/grpc/user"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func toUserSignInRequest(r *SignInRequest) *user.SignInRequest {
	return &user.SignInRequest{
		Email:    r.Email,
		Password: r.Password,
	}
}

type SignUpRequest struct {
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
}

func toUserSignUpRequest(r *SignUpRequest) *user.SignUpRequest {
	return &user.SignUpRequest{
		Email:     r.Email,
		Password:  r.Password,
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Phone:     r.Phone,
	}
}
