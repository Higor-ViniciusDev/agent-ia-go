// internal/app/app.go
package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
	services "github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type App struct {
	Config *config.Config
}

func New(cfg *config.Config) *App {
	return &App{Config: cfg}
}

func (a *App) Run(ctx context.Context) error {
	// gRPC
	grpcServer := grpc.NewServer()
	helloService := services.NewHelloService("hello clean")

	pb.RegisterHelloWorldServer(grpcServer, helloService)

	lis, err := net.Listen("tcp", ":"+a.Config.GRPCPort)
	if err != nil {
		return err
	}

	go func() {
		fmt.Println("gRPC rodando na porta", a.Config.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	// HTTP Gateway
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterHelloWorldHandlerFromEndpoint(
		ctx,
		mux,
		"localhost:"+a.Config.GRPCPort,
		opts,
	)
	if err != nil {
		return err
	}

	fmt.Println("HTTP rodando na porta", a.Config.WebPort)
	return http.ListenAndServe(":"+a.Config.WebPort, mux)
}
