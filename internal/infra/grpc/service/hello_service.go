package services

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
)

type HelloService struct {
	Message string
	pb.UnimplementedHelloWorldServer
}

func (hc *HelloService) Hello(ctx context.Context, in *pb.Blank) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: hc.Message,
	}, nil
}

func NewHelloService(message string) *HelloService {
	return &HelloService{
		Message: message,
	}
}
