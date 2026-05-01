package app

import (
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/nats"
)

// app/bootstrap.go
func Start(cfg *config.Config) error {
	nc, err := nats.NewConnectionNats(cfg.BrokerUrl, cfg.BrokerPort)
	if err != nil {
		return err
	}

	js, err := nc.JetStream()
	if err != nil {
		return err
	}

	if err := nats.EnsureWorkStream(js); err != nil {
		return err
	}

	_, err = js.QueueSubscribe(
		"work.*",
		"workers",
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}
