package main

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/logger"
	businessUniverseRepository "auth/auth_back/pkg/repositories/businessUniverse"
	contactTypeRepository "auth/auth_back/pkg/repositories/contactType"
	customerRepository "auth/auth_back/pkg/repositories/customer"
	customerUserRepository "auth/auth_back/pkg/repositories/customerUser"
	notificationRepository "auth/auth_back/pkg/repositories/notification"
	roleRepository "auth/auth_back/pkg/repositories/role"
	userRepository "auth/auth_back/pkg/repositories/user"
	"auth/auth_back/pkg/services/grpc/user"
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

	userRepo := userRepository.ConnectRepository(db)
	roleRepo := roleRepository.ConnectRepository(db)
	customerUserRepo := customerUserRepository.ConnectRepository(db)
	businessUniverseRepo := businessUniverseRepository.ConnectRepository(db)
	contactTypeRepo := contactTypeRepository.ConnectRepository(db)
	customerRepo := customerRepository.ConnectRepository(db)
	notifyRepo := notificationRepository.ConnectRepository(db)

	listener, err := net.Listen("tcp", viper.GetString("grpc.user.host")+":"+viper.GetString("grpc.user.port"))
	// listener, err := net.Listen("tcp", ":"+viper.GetString("grpc.auth.port"))

	if err != nil {
		l.LogError("Failed to listening: "+err.Error(), "cmd/app/main.main")
		log.Fatalf("Failed to listening %v", err)
	}

	userGrpcServer := user.GrpcServer{
		UserRepo:                          userRepo,
		RoleRepo:                          roleRepo,
		CustomerUserRepo:                  customerUserRepo,
		BusinessUniverseRepo:              businessUniverseRepo,
		ContactTypeRepo:                   contactTypeRepo,
		CustomerRepo:                      customerRepo,
		NotifyRepo:                        notifyRepo,
		UnimplementedB24UserServiceServer: user.UnimplementedB24UserServiceServer{},
	}

	grpcServer := grpc.NewServer()

	user.RegisterB24UserServiceServer(grpcServer, &userGrpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		l.LogError("Failed to serve: "+err.Error(), "cmd/app/main.main")
		log.Fatalf("Failed to serve %v", err)
	}
}
