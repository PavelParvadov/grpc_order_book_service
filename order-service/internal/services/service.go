package services

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	"go.uber.org/zap"
)

type OrderService struct {
	Log *zap.Logger
	OrderSaver
	OrderProvider
}

type OrderSaver interface {
	Save(ctx context.Context, order models.Order) (string, error)
}

type OrderProvider interface {
	GetOrders(ctx context.Context) ([]models.Order, error)
}

func NewOrderService(log *zap.Logger, os OrderSaver, op OrderProvider) *OrderService {
	return &OrderService{Log: log, OrderSaver: os, OrderProvider: op}
}
