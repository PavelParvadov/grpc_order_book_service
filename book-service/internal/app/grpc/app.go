package grpc

import (
	"fmt"
	GrpcServer "github.com/PavelParvadov/grpc_order_book_service/book-service/internal/delivery/grpc"
	"github.com/PavelParvadov/grpc_order_book_service/book-service/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type BookGrpcApp struct {
	Log  *zap.Logger
	Grpc *grpc.Server
	Port int
}

func NewGRPCApp(log *zap.Logger, port int, bookService *service.BookService) *BookGrpcApp {
	server := grpc.NewServer()
	GrpcServer.RegisterGRPCServer(server, bookService)
	return &BookGrpcApp{
		Log:  log,
		Grpc: server,
		Port: port,
	}

}

func (a *BookGrpcApp) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *BookGrpcApp) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return err
	}
	a.Log.Info("grpc server listening on", zap.Int("port", a.Port))

	if err = a.Grpc.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (a *BookGrpcApp) Stop() {
	a.Log.Info("grpc server stopping")
	a.Grpc.GracefulStop()
}
