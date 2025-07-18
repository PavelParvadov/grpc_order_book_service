package main

import (
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/app"
	"github.com/PavelParvadov/grpc_order_book_service/api-gateway/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.GetConfig()

	application := app.NewApp(cfg)
	go application.HttpApp.MustRun()
	log.Println("Starting gateway")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	application.HttpApp.Stop()

}
