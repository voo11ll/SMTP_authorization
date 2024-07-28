package customerUserHandler

import "auth/auth_back/pkg/services/grpc/user"

type UserType struct {
	Id         string `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Avatar     string `json:"avatar"`
	CustomerId string `json:"customerId"`
}

type AddUserResponse struct {
	Message string   `json:"message"`
	User    UserType `json:"user"`
}

func toAddUserResponse(r *user.CustomerUserResponse) *AddUserResponse {
	return &AddUserResponse{
		Message: r.Message,
		User: UserType{
			Id:         r.CustomerUser.Id,
			FirstName:  r.CustomerUser.FirstName,
			SecondName: r.CustomerUser.SecondName,
			LastName:   r.CustomerUser.LastName,
			Email:      r.CustomerUser.Email,
			Phone:      r.CustomerUser.Phone,
			Avatar:     r.CustomerUser.Avatar,
		},
	}
}

type GetUserResponse struct {
	Code    int32    `json:"code"`
	Message string   `json:"message"`
	User    UserType `json:"user"`
}

func toUserGetResponse(r *user.CustomerUserResponse) *GetUserResponse {
	return &GetUserResponse{
		Message: r.Message,
		Code:    r.Code,
		User: UserType{
			Id:         r.CustomerUser.Id,
			FirstName:  r.CustomerUser.FirstName,
			LastName:   r.CustomerUser.LastName,
			SecondName: r.CustomerUser.SecondName,
			Email:      r.CustomerUser.Email,
			Phone:      r.CustomerUser.Phone,
			Avatar:     r.CustomerUser.Avatar,
			CustomerId: r.CustomerUser.Customer.Id,
		},
	}
}

type UpdateResponse struct {
	Message string   `json:"message"`
	Code    int32    `json:"code"`
	User    UserType `json:"user"`
}

func toUpdateResponse(r *user.CustomerUserResponse) *UpdateResponse {
	return &UpdateResponse{
		Code:    r.Code,
		Message: r.Message,
		User: UserType{
			Id:         r.CustomerUser.Id,
			FirstName:  r.CustomerUser.FirstName,
			LastName:   r.CustomerUser.LastName,
			SecondName: r.CustomerUser.SecondName,
			Email:      r.CustomerUser.Email,
			Phone:      r.CustomerUser.Phone,
			Avatar:     r.CustomerUser.Avatar,
			CustomerId: r.CustomerUser.Customer.Id,
		},
	}
}
