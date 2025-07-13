package router

import (
	"net/http"

	"github.com/example/user-service/api/handler"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.Ping)
	return mux
}
