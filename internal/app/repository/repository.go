package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	IGameRepostiory
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		IGameRepostiory: NewGameRepository(db),
	}
}
