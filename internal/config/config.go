package config

import (
	"os"

	"github.com/Higor-ViniciusDev/agent-ia-go/configuration/logger"
)

type Config struct {
	WebPort  string
	GRPCPort string
}

func Load() *Config {
	if err := logger.GetLogger().Sync(); err != nil {
		panic("logger error uninitialized")
	}

	return &Config{
		WebPort:  getEnv("WEB_SERVER_PORT", "8080"),
		GRPCPort: getEnv("GRPC_PORT", "50051"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
