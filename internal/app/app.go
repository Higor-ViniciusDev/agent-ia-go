package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/service"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Run(ctx context.Context) error {
	// --- gRPC ---
	grpcServer := grpc.NewServer()
	pb.RegisterHealthServer(grpcServer, service.NewHealthService())

	lis, err := net.Listen("tcp", ":"+a.cfg.GRPCPort)
	if err != nil {
		return fmt.Errorf("failed in heard gRPC: %w", err)
	}

	go func() {
		logger.Info("grpc has been initialized", zap.String("port", a.cfg.GRPCPort))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("gRPC ended", err)
		}
	}()

	// --- HTTP Gateway ---
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), // ← obrigatório
	}

	if err := pb.RegisterHealthHandlerFromEndpoint(ctx, mux, "localhost:"+a.cfg.GRPCPort, opts); err != nil {
		return fmt.Errorf("error in registed gateway: %w", err)
	}

	httpServer := &http.Server{
		Addr:    ":" + a.cfg.WebPort,
		Handler: mux,
	}

	// --- Graceful Shutdown ---
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("HTTP initialized", zap.String("port", a.cfg.WebPort))
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("HTTP ended with error", err)
		}
	}()

	<-quit
	logger.Info("Shutdown app...")

	grpcServer.GracefulStop()
	return httpServer.Shutdown(ctx)
}
