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

// func toBusinessUniverseType(r *user.BusinessUniverse) businessUniverseHandler.BusinessUniverseType {
// 	if r == nil {
// 		return businessUniverseHandler.BusinessUniverseType{}
// 	}

// 	return businessUniverseHandler.BusinessUniverseType{
// 		Id:     r.Id,
// 		Name:   r.Name,
// 		Domain: r.Domain,
// 	}
// }

// func toCustomerType(r *user.Customer) customerHandler.CustomerType {
// 	if r == nil {
// 		return customerHandler.CustomerType{}
// 	}

// 	return customerHandler.CustomerType{
// 		Id:               r.Id,
// 		Name:             r.Name,
// 		FullName:         r.FullName,
// 		INN:              r.Inn,
// 		KPP:              r.Kpp,
// 		LegalAddress:     r.LegalAddress,
// 		Banks:            toBanksResponse(r.Banks),
// 		ContactInfos:     toContactInfosResponse(r.ContactInfos),
// 		BusinessUniverse: customerHandler.BusinessUniverseType(toBusinessUniverseType(r.BusinessUniverse)),
// 	}
// }

// func toBankResponse(r *user.Bank) *customerHandler.BankType {
// 	return &customerHandler.BankType{
// 		Id:                r.Id,
// 		Name:              r.Name,
// 		AccountNumber:     r.AccountNumber,
// 		Bik:               r.Bik,
// 		CorrAccountNumber: r.CorrAccountNumber,
// 	}
// }

// func toBanksResponse(r []*user.Bank) []*customerHandler.BankType {
// 	var out = make([]*customerHandler.BankType, len(r))

// 	for i, _u := range r {
// 		out[i] = toBankResponse(_u)
// 	}

// 	return out
// }

// func toContactInfoResponse(r *user.ContactInfo) *customerHandler.ContactInfoType {
// 	return &customerHandler.ContactInfoType{
// 		Id: r.Id,
// 		ContactType: customerHandler.ContactType{
// 			Id:     r.ContactType.Id,
// 			Name:   r.ContactType.Name,
// 			Notion: r.ContactType.Notion,
// 		},
// 		Value: r.Value,
// 	}
// }

// func toContactInfosResponse(r []*user.ContactInfo) []*customerHandler.ContactInfoType {
// 	var out = make([]*customerHandler.ContactInfoType, len(r))

// 	for i, _u := range r {
// 		out[i] = toContactInfoResponse(_u)
// 	}

// 	return out
// }
