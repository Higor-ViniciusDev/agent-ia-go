package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Higor-ViniciusDev/agent-ia-go/configuration/logger"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/api/web"
)

func main() {
	defer func() {
		if err := logger.GetLogger().Sync(); err != nil {
			panic("logger error uninitialized")
		}
	}()

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

	webserver.InitWebServer()
}
