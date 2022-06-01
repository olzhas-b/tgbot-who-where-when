package repository

import "database/sql"

type GameRepository struct {
	DB *sql.DB
}

func (repo *GameRepository) GetListOfLeaderboard() {

}
