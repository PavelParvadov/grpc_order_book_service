package grpc

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	orderServicev1 "github.com/PavelParvadov/grpc_order_book_service/order-service/protoc/gen/go/order-service"
	"google.golang.org/grpc"
)

type OrderService interface {
	AddNewOrder(ctx context.Context, order models.Order) (string, error)
	GetOrders(ctx context.Context) ([]models.Order, error)
}

type grpcServer struct {
	orderService OrderService
	orderServicev1.UnimplementedOrderServer
}

func RegisterOrderServiceServer(server *grpc.Server, orderService OrderService) {
	orderServicev1.RegisterOrderServer(server, &grpcServer{
		orderService: orderService,
	})
}
