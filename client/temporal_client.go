package client

import "github.com/example/user-service/config"

// TemporalClient placeholder for Temporal connection.
type TemporalClient struct {
	URL string
}

func NewTemporalClient(cfg config.Config) (*TemporalClient, error) {
	return &TemporalClient{URL: cfg.TemporalURL}, nil
}
