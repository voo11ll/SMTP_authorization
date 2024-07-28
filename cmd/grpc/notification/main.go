package main

import (
	"auth/auth_back/config"
	"auth/auth_back/pkg/dbs"
	"auth/auth_back/pkg/logger"
	notificationRepository "auth/auth_back/pkg/repositories/notification"
	"auth/auth_back/pkg/services/grpc/notification"
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

	notificationRepo := notificationRepository.ConnectRepository(db)

	listener, err := net.Listen("tcp", viper.GetString("grpc.notification.host")+":"+viper.GetString("grpc.notification.port"))
	if err != nil {
		l.LogError("Failed to listening: "+err.Error(), "cmd/notification/main.main")
		log.Fatalf("Failed to listening %v", err)
	}

	grpcServer := grpc.NewServer()

	notification.RegisterNotificationServiceServer(grpcServer, &notification.GrpcServer{
		NotifyRepo:                             notificationRepo,
		UnimplementedNotificationServiceServer: notification.UnimplementedNotificationServiceServer{},
	})

	if err := grpcServer.Serve(listener); err != nil {
		l.LogError("Failed to serve: "+err.Error(), "cmd/app/main.main")
		log.Fatalf("Failed to serve %v", err)
	}
}
