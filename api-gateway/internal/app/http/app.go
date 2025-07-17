package http

import (
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/book"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/order"
	//httpDelivery "github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/delivery/http"
	"net/http"
)

type AppHttp struct {
	server *http.Server
	port   int
}

func NewAppHttp(port int, bs *book.Client, oc *order.Client) {

}
