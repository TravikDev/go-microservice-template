package main

import (
	"log"

	"github.com/example/user-service/workflow"
)

func main() {
	log.Println("Starting Temporal worker")
	if err := workflow.StartWorker(); err != nil {
		log.Fatal(err)
	}
}
