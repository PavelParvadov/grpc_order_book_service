package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/book"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/clients/order"
	httpDelivery "github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/delivery/http"
)

type AppHttp struct {
	server *http.Server
}

func NewAppHttp(port int, bs *book.Client, oc *order.Client) *AppHttp {
	handler := httpDelivery.NewHandler(bs, oc)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books", handler.GetBooks)
	mux.HandleFunc("POST /book", handler.AddBook)
	mux.HandleFunc("GET /orders", handler.GetOrders)
	mux.HandleFunc("POST /order", handler.AddOrder)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	return &AppHttp{
		server: &srv,
	}

}

func (a *AppHttp) MustRun() {
	if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
}

func (a *AppHttp) Start() error {
	return a.server.ListenAndServe()
}

func (a *AppHttp) Stop() error {
	return a.server.Shutdown(context.Background())
}
