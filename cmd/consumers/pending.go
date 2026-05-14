package main

import (
	"fmt"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/consumer"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/nats"
	"github.com/joho/godotenv"
)

func main() {
	//cod inative for production
	_ = godotenv.Load()
	// ctx := context.Background()
	cfg := config.Load()

	conNats, err := nats.NewConnectionNats(cfg.BrokerUrl, cfg.BrokerPort)
	if err != nil {
		panic(fmt.Errorf("failed load nats broker: %w", err))
	}
	defer conNats.Close()

	js, err := conNats.JetStream()
	if err != nil {
		panic(err)
	}

	if err := nats.EnsureWorkStream(js); err != nil {
		panic(fmt.Errorf("Error in create fila WORKS: %w", err))
	}

	consumeWork := consumer.NewWorkPendingConsumer(js)

	go func() {
		if err := consumeWork.Start(); err != nil {
			panic(fmt.Errorf("failed to start consumer: %w", err))
		}
	}()
	select {}
}
