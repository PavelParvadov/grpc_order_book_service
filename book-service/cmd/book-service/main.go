package main

import (
	"github.com/PavelParvadov/grpc_order_book_service/book-service/pkg/logging"
	"go.uber.org/zap"
)

func main() {
	//cfg := config.GetInstance()
	logger := logging.GetLogger()
	logger.Log(zap.InfoLevel, "инфо")
}
