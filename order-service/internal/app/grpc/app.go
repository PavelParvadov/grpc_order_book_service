package grpc

import (
	"fmt"
	grpcServer "github.com/PavelParvadov/grpc_order_book_service/order-service/internal/delivery/grpc"
	"github.com/PavelParvadov/grpc_order_book_service/order-service/internal/services"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type AppGrpc struct {
	GrpcServ *grpc.Server
	Log      *zap.Logger
	Port     int
}

func NewAppGrpc(log *zap.Logger, port int, orderService *services.OrderService) *AppGrpc {
	server := grpc.NewServer()
	grpcServer.RegisterOrderServiceServer(server, orderService)
	return &AppGrpc{
		GrpcServ: server,
		Log:      log,
		Port:     port,
	}
}
func (a *AppGrpc) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *AppGrpc) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.Port))
	if err != nil {
		return err
	}
	a.Log.Info("grpc server listening on", zap.Int("port", a.Port))
	if err := a.GrpcServ.Serve(lis); err != nil {
		return err
	}
	return nil

}

func (a *AppGrpc) Stop() {
	a.Log.Info("grpc server stopped")
	a.GrpcServ.GracefulStop()
}
