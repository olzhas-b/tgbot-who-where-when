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
	dns := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", config.Database.Host, config.Database.User, config.Database.Name, config.Database.Password, config.Database.Port)
	dns = "postgres://prgsupezuxwnyl:37e3acc567a1544135dc73692490c626b7dfeb885d74268eba1fcde46e91e8a3@ec2-34-248-169-69.eu-west-1.compute.amazonaws.com:5432/demo702cgr2ogu"
	conn, _ := pgxpool.Connect(ctx, dns)
	if err := conn.Ping(ctx); err != nil {
		log.Fatal("error pinging db: ", err)
	}
	return conn
}
