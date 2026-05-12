package consumer

import (
	"encoding/json"

	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/nats-io/nats.go"
)

type WorkPendingConsumer struct {
	channel *nats.Conn
}
type workPendingMessage struct {
	Id string `json:"id"`
}

func NewEmailPendingConsumer(channel *nats.Conn) *WorkPendingConsumer {
	return &WorkPendingConsumer{
		channel: channel,
	}
}

func (c *WorkPendingConsumer) Start(js nats.JetStreamContext) error {
	_, err := js.QueueSubscribe(
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

	logger.Info("WorkPendingConsumer: received work id " + payload.Id)
}
