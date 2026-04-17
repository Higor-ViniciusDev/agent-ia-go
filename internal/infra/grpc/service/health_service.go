package service

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
)

type HealthService struct {
	*pb.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (h *HealthService) Health(context.Context, *pb.Blank) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{
		Status: "ok",
	}, nil
}
