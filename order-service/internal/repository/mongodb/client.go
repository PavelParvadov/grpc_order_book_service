package mongodb

import (
	"context"
	"fmt"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Storage struct {
	Db *mongo.Client
}

func NewStorage(cfg *config.Config) *Storage {
	connStr := fmt.Sprintf("mongodb://%s:%s/", cfg.DBConfig.Host, cfg.DBConfig.Port)
	opts := options.Client().ApplyURI(connStr)
	if cfg.DBConfig.Username != "" && cfg.DBConfig.Password != "" {
		opts = opts.SetAuth(options.Credential{
			Username: cfg.DBConfig.Username,
			Password: cfg.DBConfig.Password,
		})
	}
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	return &Storage{
		Db: client,
	}
}
