package main

import (
	"auth/auth_back/config"
	auth "auth/auth_back/pkg/services/grpc/user"
	"context"
	"log"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO: dfdfdfd lkfsdgjldfjgldfkgj llkfdlkjgdf

func main() {
	testSignUpUser()
	testSignInUser()
	testBUCreate()
	testBUFind()
	// testRoleNameFind()
	testRolesFind()
	testRoleFind()
}

func testSignUpUser() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.auth.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := auth.NewUserServiceClient(conn)

	response, err := c.SignUp(context.Background(), &auth.SignUpRequest{
		Email:     "eddy.play@mail.ru",
		Password:  "123456",
		Phone:     "1234567",
		FirstName: "User",
		LastName:  "Test",
	})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	log.Printf("Response from server: %s %s", response.User.Id, response.Message)
}

func testSignInUser() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.auth.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := auth.NewUserServiceClient(conn)

	response, err := c.SignIn(context.Background(), &auth.SignInRequest{
		Email:    "eddy.play@mail.ru",
		Password: "123456",
	})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	// log.Printf("Response from server: %s %d %s", response.User.Id, response.Code, response.Message)
	log.Printf("Response from server: %s", response)
}

func testBUCreate() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.organization.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

}

func testBUFind() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.universe_server_port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

}

func testRolesFind() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.role_server_port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := auth.NewUserServiceClient(conn)

	response, err := c.GetRoles(context.Background(), &auth.RolesGetRequest{})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	log.Printf("Response from server: %d %s %s", response.Code, response.Message, response.Roles)
}

func testRoleFind() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+viper.GetString("grpc.role_server_port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := auth.NewUserServiceClient(conn)

	response, err := c.GetRole(context.Background(), &auth.RoleGetRequest{
		Id: "6b6634c1-364d-4dc7-a999-9ab2688a60b1",
	})

	if err != nil {
		log.Fatalf("Error when calling creating user: %s", err)
	}

	log.Printf("Response from server: %d %s %s %s", response.Code, response.Message, response.Role.Id, response.Role.Name)
}
