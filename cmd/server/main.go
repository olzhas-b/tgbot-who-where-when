package main

import (
	"fmt"
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/service"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		log.Fatalf("server.cfg.InitConfig got err %v", err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.HTTP.Name, cfg.HTTP.Port))
	if err != nil {
		log.Fatalf("server.cfg.Listen got err %v", err)
	}

	grpcServer := grpc.NewServer()
	server := service.GameServiceServer{}
	api.RegisterGameServer(grpcServer, &server)

	log.Println("trying to run grpcServer")
	log.Fatal("grpc server got error ", grpcServer.Serve(listener))
}
