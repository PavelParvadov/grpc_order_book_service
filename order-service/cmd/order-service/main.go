package main

import (
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/app"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/config"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/pkg/logging"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	_ = godotenv.Load()

	cfg := config.GetInstance()

	logger := logging.NewLogger()
	logger.Info("Инфо", zap.Any("cfg", cfg))

	application := app.NewApp(cfg, logger)
	go application.AppGrpc.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	application.Stop()

}
