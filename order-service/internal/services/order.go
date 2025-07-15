package services

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	"go.uber.org/zap"
)

func (s *OrderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	orders, err := s.GetOrders(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *OrderService) AddNewOrder(ctx context.Context, order models.Order) (string, error) {
	_, err := s.BookClient.GetBookByID(ctx, order.BookId)
	if err != nil {
		s.Log.Warn("book not found", zap.Error(err))
		return "", ErrBookNotFound
	}
	id, err := s.Save(ctx, order)
	if err != nil {
		s.Log.Warn("cannot save order", zap.Error(err))
		return "", err
	}
	s.Log.Info("new order", zap.String("id", id))
	return id, nil
}
