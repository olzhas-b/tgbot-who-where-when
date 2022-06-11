package main

import (
	"fmt"
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/repository"
	"gitlab.ozon.dev/hw/homework-2/internal/app/service"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	db2 "gitlab.ozon.dev/hw/homework-2/internal/db"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("server.cfg.InitConfig got err %v", err)
	}
	grpcAddress := fmt.Sprintf(":%s", cfg.GRPC.Port)
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("server.cfg.Listen got err %v", err)
	} else {
		log.Println("tcp was listen successfully: ", grpcAddress)
	}
	db := db2.InitDB(cfg)
	repo := repository.New(db)
	server := service.NewGameServiceServer(repo)

	grpcServer := grpc.NewServer()
	api.RegisterGameServer(grpcServer, server)

	log.Println("trying to run grpcServer")
	log.Fatal("grpc server got error ", grpcServer.Serve(listener))
}
