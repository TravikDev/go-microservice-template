package router

import (
	"net/http"
  
	"github.com/example/user-service/api/handler"
	"github.com/example/user-service/service"
)

func New(svc *service.UserService) http.Handler {
	mux := http.NewServeMux()
	userHandler := handler.NewUserHandler(svc)
	charHandler := handler.NewCharacterHandler(svc)

	mux.HandleFunc("/ping", handler.Ping)
	mux.HandleFunc("/users", userHandler.ListUsers)
	mux.HandleFunc("/users/", charHandler.ListCharacters)
}

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	return mux
}
