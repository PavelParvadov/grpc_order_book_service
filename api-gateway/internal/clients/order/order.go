package order

import (
	"context"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/dto"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/domain/models"
	orderServicev1 "github.com/PavelParvadov/grpc_order_book_service/order-service/protoc/gen/go/order-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type Client struct {
	Client orderServicev1.OrderClient
}

func NewClient(ctx context.Context, addr string) (*Client, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{orderServicev1.NewOrderClient(cc)}, nil
}

func (c *Client) AddOrder(ctx context.Context, in dto.AddOrderInput) (dto.AddOrderOutput, error) {
	order, err := c.Client.AddOrder(ctx, &orderServicev1.AddOrderRequest{
		BookId: in.BookId,
		Place:  in.Place,
		Price:  in.Price,
		Status: toOrderStatusEnum(in.Status),
	})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return dto.AddOrderOutput{}, ErrInternalServerError
		}
		return dto.AddOrderOutput{}, err
	}
	return dto.AddOrderOutput{
		ID: order.OrderId,
	}, nil
}

func (c *Client) GetOrders(ctx context.Context) ([]models.Order, error) {
	resp, err := c.Client.GetOrders(ctx, &orderServicev1.GetOrdersRequest{})
	if err != nil {
		_, ok := status.FromError(err)
		if ok {
			return []models.Order{}, ErrInternalServerError
		}
		return []models.Order{}, err
	}

	orders := make([]models.Order, 0, len(resp.GetOrderData()))
	for _, order := range resp.GetOrderData() {
		orders = append(orders, models.Order{
			ID:     order.OrderId,
			BookID: order.BookId,
			Place:  order.Place,
			Price:  order.Price,
			Status: order.Status.String(),
		})
	}
	return orders, nil
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
