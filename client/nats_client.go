package client

import "github.com/example/user-service/config"

// NatsClient placeholder for NATS connection.
type NatsClient struct {
	URL string
}

func NewNatsClient(cfg config.Config) (*NatsClient, error) {
	return &NatsClient{URL: cfg.NatsURL}, nil
}
