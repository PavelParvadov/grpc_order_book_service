package app

import (
	grpcapp "github.com/PavelParvadov/grpc_order_book_service/book-service/internal/app/grpc"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/config"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/repository/postgres"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/service"
	"go.uber.org/zap"
)

type App struct {
	BookGrpc *grpcapp.BookGrpcApp
}

func NewApp(log *zap.Logger, cfg config.Config) *App {
	storage := postgres.NewStorage(&cfg)
	BookService := service.NewBookService(log, storage, storage)
	GrpcServer := grpcapp.NewGRPCApp(log, cfg.GRPCConf.Port, BookService)
	return &App{
		BookGrpc: GrpcServer,
	}
}

func (a *App) Stop() {
	a.BookGrpc.Stop()
}
