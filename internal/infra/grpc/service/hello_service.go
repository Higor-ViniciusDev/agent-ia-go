package services

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
)

type HelloService struct {
	pb.UnimplementedHelloWorldServer
}

func (hc *HelloService) Hello(ctx context.Context, in *pb.Blank) (*pb.Hello, error) {
	return &pb.Hello{
		Message: "Ola mundo",
	}, nil
}

func NewOrderService() *HelloService {
	return &HelloService{}
}
