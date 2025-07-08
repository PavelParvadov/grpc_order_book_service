package postgres

import (
	"context"
	"fmt"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Storage struct {
	DbPool *pgxpool.Pool
}

func NewStorage(cfg *config.Config) *Storage {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DbConf.Host,
		cfg.DbConf.Port,
		cfg.DbConf.Username,
		cfg.DbConf.Password,
		cfg.DbConf.Dbname)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		panic(err)
	}
	return &Storage{
		DbPool: pool,
	}
}
