package consumer

import (
	"context"
	"encoding/json"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/repository"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/nats-io/nats.go"
)

type WorkPendingConsumer struct {
	channel nats.JetStreamContext
	repo    *repository.WorkRepository
}
type workPendingMessage struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func NewWorkPendingConsumer(channel nats.JetStreamContext, repository *repository.WorkRepository) *WorkPendingConsumer {
	return &WorkPendingConsumer{
		channel: channel,
		repo:    repository,
	}
}

func (c *WorkPendingConsumer) Start(ctx context.Context) error {
	logger.Info("Starting Consumer pending")

	_, err := c.channel.QueueSubscribe(
		"work.pending",
		"workers",
		func(msg *nats.Msg) {
			c.handle(msg, ctx)
			msg.Ack()
		},
		nats.Durable("workers"),
		nats.ManualAck(),
	)

	return err
}

func (c *WorkPendingConsumer) handle(msg *nats.Msg, ctx context.Context) {
	var payload workPendingMessage
	if err := json.Unmarshal(msg.Data, &payload); err != nil {
		logger.Error("WorkPendingConsumer: invalid payload", err)
		msg.Nak()
		return
	}

	c.repo.UpdateStatus(ctx, payload.Id, entity.WorkStatusProcessing)
}
