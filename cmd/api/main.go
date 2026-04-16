package main

import (
	"context"
	"log"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/app"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	application := app.New(cfg)

	if err := application.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
