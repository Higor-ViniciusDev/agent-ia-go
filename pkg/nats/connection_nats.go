package nats

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func NewConnectionNats(url string, port string) (*nats.Conn, error) {
	nc, err := nats.Connect(fmt.Sprintf("nats://%v:%v", url, port),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)

	return nc, err
}

// infra/nats_stream.go
func EnsureWorkStream(js nats.JetStreamContext) error {
	_, err := js.StreamInfo("WORK")
	if err == nil {
		return nil
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "WORK",
		Subjects: []string{"work.*"},
	})
	return err
}
