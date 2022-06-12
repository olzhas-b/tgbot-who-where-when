package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/hw/homework-2/internal/config"
	"log"
)

func InitDB(config config.Config) *pgxpool.Pool {
	ctx := context.Background()
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	println("dns", dns)
	conn, err := pgxpool.Connect(ctx, dns)
	if err != nil {
		log.Fatal("error to connect: ", err)
	}
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}
	return conn
}
