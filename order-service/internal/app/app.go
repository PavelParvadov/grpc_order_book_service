package app

import (
	"context"
	bookservice "github.com/PavelParvadov/grpc_order_book_service/order-service/clients/book-service"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/app/grpc"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/config"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/repository/mongodb"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/services"
	"go.uber.org/zap"
	"time"
)

type App struct {
	AppGrpc *grpc.AppGrpc
}

func NewApp(cfg *config.Config, log *zap.Logger) *App {
	storage := mongodb.NewStorage(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bookClient, err := bookservice.NewBookClient(ctx, "localhost:5555")
	if err != nil {
		panic(err)
	}
	OrderService := services.NewOrderService(log, storage, storage, bookClient)
	server := grpc.NewAppGrpc(log, cfg.GrpcConfig.Port, OrderService)
	return &App{
		AppGrpc: server,
	}
}

func (a *App) Stop() {
	a.AppGrpc.Stop()
}
