package config

import "os"

// Config holds application configuration.
type Config struct {
	NatsURL     string
	TemporalURL string
	DBURL       string
}

// Load reads configuration from environment variables.
func Load() Config {
	return Config{
		NatsURL:     os.Getenv("NATS_URL"),
		TemporalURL: os.Getenv("TEMPORAL_URL"),
		DBURL:       os.Getenv("DATABASE_URL"),
	}
}
