package user

import (
	"auth/auth_back/config"
	modelsOrganization "auth/auth_back/models/organization"
	modelsUser "auth/auth_back/models/user"
	"auth/auth_back/pkg/services/grpc/organization"

	"context"
	"log"

	customerUserRepository "auth/auth_back/pkg/repositories/customerUser"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetCustomerGrpcClient(id string) *organization.CustomerResponse {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	var response *organization.CustomerResponse

	response, err = c.GetCustomer(context.Background(), &organization.GetRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	return response
}

func toCustomerUserResponse(in *modelsOrganization.CustomerUser, role *modelsUser.Role, businessUniverse *modelsOrganization.BusinessUniverse, customer *modelsOrganization.Customer, s *GrpcServer) *CustomerUser {
	return &CustomerUser{
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
		BusinessUniverse: &BusinessUniverse{
			Id:     businessUniverse.ID.String(),
			Name:   businessUniverse.Name,
			Domain: businessUniverse.Domain,
		},
		Customer: &Customer{
			Id:           customer.ID.String(),
			Name:         customer.Name,
			FullName:     customer.FullName,
			Inn:          customer.INN,
			Kpp:          customer.KPP,
			LegalAddress: customer.LegalAddress,
			Banks:        toBanksGet(customer.Banks),
			ContactInfos: toCompanyContactInfosGet(customer.ContactInfos, s),
		},
	}
}

func toCustomerCreateUser(u *CustomerUserAddRequest) *customerUserRepository.User {
	return &customerUserRepository.User{
		Email:              u.GetEmail(),
		Password:           u.GetPassword(),
		BusinessUniverseID: u.GetBusinessUniverseID(),
		CustomerID:         u.GetCustomerID(),
		RoleID:             u.GetRoleID(),
	}
}

func toCustomerUserUpdate(u *CustomerUserUpdateRequest) *customerUserRepository.User {
	return &customerUserRepository.User{
		FirstName:  u.GetFirstName(),
		LastName:   u.GetLastName(),
		SecondName: u.GetSecondName(),
		Phone:      u.GetPhone(),
	}
}

func toBanksGet(u []*modelsOrganization.CustomerBank) []*Bank {
	var out = make([]*Bank, len(u))

	for i, _u := range u {
		out[i] = &Bank{
			Id:                _u.ID.String(),
			Name:              _u.Name,
			AccountNumber:     _u.AccountNumber,
			Bik:               _u.Bik,
			CorrAccountNumber: _u.CorrAccountNumber,
			IsOpen:            _u.IsOpen,
		}
	}

	return out
}

func toCompanyContactInfosGet(u []*modelsOrganization.CustomerContactInfo, s *GrpcServer) []*ContactInfo {
	var out = make([]*ContactInfo, len(u))
	var ctx context.Context

	for i, _u := range u {
		contactType := s.ContactTypeRepo.FindItemById(ctx, _u.ContactTypeID)
		out[i] = &ContactInfo{
			Id: _u.ID.String(),
			ContactType: &ContactType{
				Id:     contactType.ID.String(),
				Name:   contactType.Name,
				Notion: contactType.Notion,
			},
			Value: _u.Value,
		}
	}

	return out
}
