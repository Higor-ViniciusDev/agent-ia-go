package handlers

import (
	"encoding/json"
	"sync"

	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/events"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/nats-io/nats.go"
)

type CreateWorkHandler struct {
	BrokerChannel *nats.Conn
}

func NewWorkCreatedHandler(channel *nats.Conn) *CreateWorkHandler {
	return &CreateWorkHandler{
		BrokerChannel: channel,
	}
}

func (ch *CreateWorkHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()

	jsonOutput, err := json.Marshal(event.GetPayload())

	if err != nil {
		logger.Error("Error in convert json, date paloyd its not valid? ", err)
		jsonOutput = []byte("{error:true}")
	}

	err = ch.BrokerChannel.Publish("work.create", jsonOutput)

	if err != nil {
		logger.Error("Publish create work failed, error in comunnication with broker", err)
	}
}
