package subscriber

import (
	"log"

	"github.com/example/user-service/client"
	"github.com/example/user-service/config"
)

func Start() error {
	cfg := config.Load()
	_, err := client.NewNatsClient(cfg)
	if err != nil {
		return err
	}
	log.Println("Subscribed to user events")
	return nil
}
