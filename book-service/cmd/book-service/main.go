package main

import (
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/app"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/config"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/pkg/logging"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetInstance()
	logger := logging.GetLogger()
	logger.Log(zap.InfoLevel, "инфо")

	application := app.NewApp(logger, *cfg)
	go application.BookGrpc.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	logger.Info("Shutting down server...")
	application.Stop()

}
