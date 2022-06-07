package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/hw/homework-2/internal/app/models"
	"log"
)

type GameRepository struct {
	db *pgxpool.Pool
}

func NewGameRepository(db *pgxpool.Pool) *GameRepository {
	return &GameRepository{db: db}
}

func (repo *GameRepository) GetLeaderboards(ctx context.Context, chatID int64) (users []models.User, err error) {
	query := `SELECT full_name, user_name, score
				FROM leader_board WHERE chat_id = $1
				ORDER BY score DESC
				LIMIT 10;`
	if err := pgxscan.Select(ctx, repo.db, &users, query, chatID); err != nil {
		log.Printf("GameRepository.GetLeaderboards got error %v", users)
	}
	return
}

func (repo *GameRepository) GetRandomQuestion(ctx context.Context) (question, answer string, err error) {
	if err := repo.db.QueryRow(ctx, "SELECT text, answer FROM question ORDER BY RANDOM() LIMIT 1;").
		Scan(&question, &answer); err != nil {
		return "", "", err
	}
	return
}

func (repo *GameRepository) InsertQuestion(ctx context.Context, question, answer string) error {
	query := `INSERT INTO question(text, answer)
				SELECT $1, $2
				WHERE NOT EXISTS
					(SELECT id FROM question WHERE text = $3 and answer = $4);`
	_, err := repo.db.Exec(ctx, query, question, answer, question, answer)
	return err
}

func (repo *GameRepository) AddPointToUser(ctx context.Context, req models.Request) error {
	tx, err := repo.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()
	_, err = tx.Exec(ctx, `
			INSERT INTO leader_board(chat_id, user_id, user_name, full_name)
			SELECT $1, $2, $3, $4
			WHERE NOT EXISTS (SELECT chat_id FROM leader_board WHERE chat_id = $5 and user_id = $6);`,
		req.ChatID, req.UserID, req.UserName, req.FullName, req.ChatID, req.UserID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx,
		"UPDATE leader_board SET score=score+1 "+
			"WHERE chat_id = $1 and user_id = $2;", req.ChatID, req.UserID)

	return err
}
