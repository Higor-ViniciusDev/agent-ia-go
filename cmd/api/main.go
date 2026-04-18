package main

import (
	"context"
	"log"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/app"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
)

func main() {
	//cod inative for production
	// _ = godotenv.Load()
	ctx := context.Background()
	cfg := config.Load()

	if err := app.New(cfg).Run(ctx); err != nil {
		log.Fatal(err)
	}
}
