package main

import (
	"log"

	"github.com/example/user-service/subscriber"
)

func main() {
	log.Println("Starting NATS subscriber")
	if err := subscriber.Start(); err != nil {
		log.Fatal(err)
	}
}
