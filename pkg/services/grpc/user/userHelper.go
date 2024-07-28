package user

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/services/grpc/organization"

	"context"
	"log"

	userRepository "auth/auth_back/pkg/repositories/user"

	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetAndCreateBusinessUniverseGrpcClient(id string) *organization.BusinessUniverseGetResponse {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	conn, err := grpc.Dial(viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := organization.NewB24OrganizationServiceClient(conn)

	var response *organization.BusinessUniverseGetResponse
	var resp *organization.BusinessUniverseCreateResponse

	if id == "" {
		resp, err = c.CreateBusinessUniverse(context.Background(), &organization.BusinessUniverseCreateRequest{})

		if err != nil {
			log.Fatalf("Error when calling creating BU: %s", err)
		}

		id = resp.BusinessUniverse.Id
	}

	response, err = c.GetBusinessUniverse(context.Background(), &organization.BusinessUniverseGetRequest{
		Id: id,
	})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	return response
}

func toSignUp(u *SignUpRequest) *userRepository.User {
	return &userRepository.User{
		Email:     u.GetEmail(),
		Password:  u.GetPassword(),
		FirstName: u.GetFirstName(),
		LastName:  u.GetLastName(),
		Phone:     u.GetPhone(),
	}
}

func toUserUpdate(u *UserUpdateRequest) *userRepository.User {
	return &userRepository.User{
		FirstName:  u.GetFirstName(),
		LastName:   u.GetLastName(),
		SecondName: u.GetSecondName(),
		Phone:      u.GetPhone(),
	}
}
