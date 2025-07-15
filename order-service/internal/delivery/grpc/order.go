package grpc

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/domain/models"
	orderServicev1 "github.com/PavelParvadov/grpc_order_book_service/order-service/protoc/gen/go/order-service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *grpcServer) AddOrder(ctx context.Context, req *orderServicev1.AddOrderRequest) (*orderServicev1.AddOrderResponse, error) {
	order, err := validateInputData(req)
	if err != nil {
		return nil, err
	}
	id, err := s.orderService.AddNewOrder(ctx, order)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &orderServicev1.AddOrderResponse{
		OrderId: id,
	}, nil
}

func (s *grpcServer) GetOrders(ctx context.Context, req *orderServicev1.GetOrdersRequest) (*orderServicev1.GetOrdersResponse, error) {
	orders, err := s.orderService.GetOrders(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	pbBooks := make([]*orderServicev1.OrderData, len(orders))
	for _, order := range orders {
		pbBook := &orderServicev1.OrderData{
			OrderId: order.Id.Hex(),
			BookId:  order.BookId,
			Status:  toOrderStatusEnum(order.Status),
			Price:   order.Price,
			Place:   order.Place,
		}
		pbBooks = append(pbBooks, pbBook)
	}
	return &orderServicev1.GetOrdersResponse{
		OrderData: pbBooks,
	}, nil
}

func validateInputData(req *orderServicev1.AddOrderRequest) (models.Order, error) {
	if req.BookId == 0 || req.Place == "" || req.Price == 0 || req.Status.String() == "" {
		return models.Order{}, status.Error(codes.InvalidArgument, "Invalid input")
	}
	return models.Order{
		BookId: req.BookId,
		Place:  req.Place,
		Price:  req.Price,
		Status: req.Status.String(),
	}, nil
}

func toOrderStatusEnum(status string) orderServicev1.OrderStatus {
	switch status {
	case "PENDING":
		return orderServicev1.OrderStatus_PENDING
	case "COMPLETED":
		return orderServicev1.OrderStatus_COMPLETED
	case "CANCELLED":
		return orderServicev1.OrderStatus_CANCELLED
	default:
		return orderServicev1.OrderStatus_UNKNOWN
	}
}
