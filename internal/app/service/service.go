package service

import (
	"context"
	"gitlab.ozon.dev/hw/homework-2/api"
	"gitlab.ozon.dev/hw/homework-2/internal/app/repository"
)

type GameServiceServer struct {
	repo *repository.GameRepository
	api.UnimplementedGameServer
}

func New(repo *repository.GameRepository) *GameServiceServer {
	return &GameServiceServer{repo: repo}
}

func (srv *GameServiceServer) Start(context.Context, *api.Message) (*api.Result, error) {
	return &api.Result{Text: "hello from server"}, nil
}
