package main

import (
	"log"
	"net/http"
  
	"github.com/example/user-service/api/router"
	"github.com/example/user-service/service"
)

func main() {
	svc := service.NewUserService()
	r := router.New(svc)
}

func main() {
	r := router.New()

	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
