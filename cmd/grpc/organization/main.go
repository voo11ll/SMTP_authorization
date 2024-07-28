package main

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/logger"
	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
	companyRepository "auth/auth_back/pkg/repositories/company"
	contactTypeRepository "auth/auth_back/pkg/repositories/contactType"
	customerRepository "auth/auth_back/pkg/repositories/customer"
	org "auth/auth_back/pkg/services/grpc/organization"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var l = logger.Logger{}

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	db := dbs.InitDB()

	postgres, _ := db.DB()

	defer postgres.Close()

	businessUniverseRepo := businessUniverseRepository.ConnectRepository(db)
	contactTypeRepo := contactTypeRepository.ConnectRepository(db)
	companyRepo := companyRepository.ConnectRepository(db)
	customerRepo := customerRepository.ConnectRepository(db)

	listener, err := net.Listen("tcp", viper.GetString("grpc.organization.host")+":"+viper.GetString("grpc.organization.port"))
	// listener, err := net.Listen("tcp", ":"+viper.GetString("grpc.organization.port"))

	if err != nil {
		l.LogError("Failed to listening: "+err.Error(), "cmd/app/main.main")
		log.Fatalf("Failed to listening %v", err)
	}

	orgGrpcServer := org.GrpcServer{
		BusinessUniverseRepo: businessUniverseRepo,
		ContactTypeRepo:      contactTypeRepo,
		CompanyRepo:          companyRepo,
		CustomerRepo:         customerRepo,
		UnimplementedB24OrganizationServiceServer: org.UnimplementedB24OrganizationServiceServer{},
	}

	grpcServer := grpc.NewServer()

	org.RegisterB24OrganizationServiceServer(grpcServer, &orgGrpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		l.LogError("Failed to serve: "+err.Error(), "cmd/role/main.main")
		log.Fatalf("Failed to serve %v", err)
	}
}
