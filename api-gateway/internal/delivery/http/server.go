package http

import (
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/book"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/order"
)

type Handler struct {
	BookHandler      *book.Client
	OrderBookHandler *order.Client
}

func NewHandler(bh *book.Client, obc *order.Client) *Handler {
	return &Handler{
		BookHandler:      bh,
		OrderBookHandler: obc,
	}
}
