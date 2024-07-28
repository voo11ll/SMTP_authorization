package userHandler

import "auth/auth_back/pkg/services/grpc/user"

type UserType struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
}

type GetUserResponse struct {
	Code    int32    `json:"code"`
	Message string   `json:"message"`
	User    UserType `json:"user"`
}

func toUserGetResponse(r *user.UserResponse) *GetUserResponse {
	return &GetUserResponse{
		Message: r.Message,
		Code:    r.Code,
		User: UserType{
			FirstName:  r.User.FirstName,
			LastName:   r.User.LastName,
			SecondName: r.User.SecondName,
			Email:      r.User.Email,
			Phone:      r.User.Phone,
			Avatar:     r.User.Avatar,
		},
	}
}

type UpdateResponse struct {
	Message string   `json:"message"`
	Code    int32    `json:"code"`
	User    UserType `json:"user"`
}

func toUpdateResponse(r *user.UserResponse) *UpdateResponse {
	return &UpdateResponse{
		Code:    r.Code,
		Message: r.Message,
		User: UserType{
			FirstName:  r.User.FirstName,
			LastName:   r.User.LastName,
			SecondName: r.User.SecondName,
			Email:      r.User.Email,
			Phone:      r.User.Phone,
			Avatar:     r.User.Avatar,
		},
	}
}
