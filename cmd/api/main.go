package main

import (
	"log"
	"net/http"

	"github.com/example/user-service/api/router"
	"github.com/example/user-service/client"
	"github.com/example/user-service/config"
	"github.com/example/user-service/repository"
	"github.com/example/user-service/service"
)

func main() {
	cfg := config.Load()
	db, err := client.DB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPgUserRepository(db)
	svc := service.NewUserService(repo)
	r := router.New(svc)

	log.Println("Starting HTTP server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
