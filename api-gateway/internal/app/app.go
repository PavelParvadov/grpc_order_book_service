package app

import (
	"context"
	"fmt"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/app/http"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/book"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/order"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/config"
	"time"
)

type App struct {
	HttpApp *http.AppHttp
}

func NewApp(cfg *config.Config) *App {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	bookServiceClient, err := book.NewClient(ctx, fmt.Sprintf("%s:%d", cfg.BookService.Host, cfg.BookService.Port))
	if err != nil {
		panic(err)
	}
	orderServiceClient, err := order.NewClient(ctx, fmt.Sprintf("%s:%d", cfg.OrderService.Host, cfg.OrderService.Port))
	if err != nil {
		panic(err)
	}
	httpServer := http.NewAppHttp(cfg.HttpServer.Port, bookServiceClient, orderServiceClient)
	return &App{
		HttpApp: httpServer,
	}
}
