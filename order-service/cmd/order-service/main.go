package main

import (
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/config"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/pkg/logging"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := config.GetInstance()
	logger := logging.NewLogger()
	logger.Info("Инфо", zap.Any("cfg", cfg))
}
