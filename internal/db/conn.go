package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"log"
)

func InitDB(config config.Config) *sql.DB {
	ctx := context.Background()
	dns := fmt.Sprintf("user=%s dbname=%s sslmode=disable", config.Database.User, config.Database.User)

	conn, _ := pgx.Connect(ctx, dns)
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("sql Open got error: ", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("error pinging db: ", err)
	}
	return db
}
