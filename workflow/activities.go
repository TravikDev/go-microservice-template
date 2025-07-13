package workflow

import (
	"log"

	"github.com/example/user-service/client"
	"github.com/example/user-service/config"
)

func StartWorker() error {
	cfg := config.Load()
	_, err := client.NewTemporalClient(cfg)
	if err != nil {
		return err
	}
	log.Println("Temporal worker started")
	return nil
}
