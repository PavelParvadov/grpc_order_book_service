package grpc

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	grpc_health_v1.UnimplementedHealthServer
}

func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

func (h *HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}
