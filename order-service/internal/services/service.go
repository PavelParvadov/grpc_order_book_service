package services

import (
	"context"
	book_service "github.com/PavelParvadov/grpc_order_book_service/order-service/clients/book-service"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	"go.uber.org/zap"
)

type OrderService struct {
	Log *zap.Logger
	OrderSaver
	OrderProvider
	BookClient *book_service.BookClient
}

type OrderSaver interface {
	Save(ctx context.Context, order models.Order) (string, error)
}

type OrderProvider interface {
	GetOrders(ctx context.Context) ([]models.Order, error)
}

func NewOrderService(log *zap.Logger, os OrderSaver, op OrderProvider, bc *book_service.BookClient) *OrderService {
	return &OrderService{Log: log, OrderSaver: os, OrderProvider: op, BookClient: bc}
}
