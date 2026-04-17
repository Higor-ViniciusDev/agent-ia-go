package service

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkService struct {
	pb.UnimplementedWorkServer
}

func NewWorkService() *WorkService {
	return &WorkService{}
}

func (w *WorkService) WorkAction(context.Context, *pb.WorkRequest) (*pb.WorkResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method WorkAction not implemented")
}
