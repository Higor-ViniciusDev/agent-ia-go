package entity

import (
	"time"

	"github.com/google/uuid"
)

type Conversation struct {
	id      uuid.UUID
	title   string
	CreatAt time.Time
}

type ConversationEntityInterface interface {
	GetValuesEntity() string
}

func (c *Conversation) GetValuesEntity() string {
	return "teste 25"
}
