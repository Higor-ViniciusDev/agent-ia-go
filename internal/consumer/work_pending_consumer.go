package consumer

import (
	"encoding/json"

	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/nats-io/nats.go"
)

type WorkPendingConsumer struct {
	channel nats.JetStreamContext
}
type workPendingMessage struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func NewWorkPendingConsumer(channel nats.JetStreamContext) *WorkPendingConsumer {
	return &WorkPendingConsumer{
		channel: channel,
	}
}

func (c *WorkPendingConsumer) Start() error {
	logger.Info("Starting Consumer pending")

	_, err := c.channel.QueueSubscribe(
		"work.pending",
		"workers",
		func(msg *nats.Msg) {
			c.handle(msg)
			msg.Ack()
		},
		nats.Durable("workers"),
		nats.ManualAck(),
	)

	return err
}

func (c *WorkPendingConsumer) handle(msg *nats.Msg) {
	var payload workPendingMessage
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		logger.Error("WorkPendingConsumer: invalid payload", err)
		msg.Nak()
		return
	}

}
