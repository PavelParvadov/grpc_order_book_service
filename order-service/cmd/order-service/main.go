package main

import "github.com/PavelParvadov/grpc_order_book_service/order-service/pkg/logging"

func main() {
	logger := logging.NewLogger()
	logger.Info("Инфо")
}
