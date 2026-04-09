package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/Higor-ViniciusDev/agent-ia-go/configuration/logger"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/api/web"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
	services "github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	defer func() {
		if err := logger.GetLogger().Sync(); err != nil {
			panic("logger error uninitialized")
		}
	}()

	if os.Getenv("ENV") == "" {
		if err := godotenv.Load(); err != nil {
			panic("error in load variables ambient")
		}
	}
	webServerPort := os.Getenv("WEB_SERVER_PORT")

	if webServerPort == "" {
		panic("Port init not found in .env")
	}

	webserver := web.NewWebServer(fmt.Sprintf(":%v", webServerPort))

	webserver.RegisterRoute("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"okay"}`))
		if err != nil {
			logger.Info("error in write message handler")
		}
	}, "get")

	webserver.RegisterRoute("/entities", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"okay"}`))
		if err != nil {
			logger.Info("error in write message handler")
		}
	}, "get")

	go webserver.InitWebServer()

	serviceGRPC := services.NewHelloService("hello teste 25")
	grpcServe := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServe, serviceGRPC)
	reflection.Register(grpcServe)

	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	fmt.Println("Servidor GRPC Rodando na porta", 50051)

	if err != nil {
		panic(err)
	}

	go grpcServe.Serve(listen)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	pb.RegisterHelloWorldHandlerFromEndpoint(ctx, mux, "app:50051", opts)

	http.ListenAndServe(":8080", mux)
}
