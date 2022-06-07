package repository

import (
	"context"
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
)

type IGameRepostiory interface {
	GetRandomQuestion(ctx context.Context) (string, string, error)
	InsertQuestion(ctx context.Context, question, answer string) error
	AddPointToUser(ctx context.Context, request models.Request) error
	GetLeaderboards(ctx context.Context, chatID int64) (users []models.User, err error)
}
