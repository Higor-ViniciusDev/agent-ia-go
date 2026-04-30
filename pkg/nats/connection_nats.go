package nats

import "github.com/nats-io/nats.go"

func NewConnectionNats() (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)

	return nc, err
}
