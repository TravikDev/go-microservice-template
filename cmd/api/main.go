package main

import (
	"log"
	"net/http"

	"github.com/example/user-service/api/router"
)

func main() {
	r := router.New()
	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
