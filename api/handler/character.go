package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/example/user-service/service"
)

// CharacterHandler handles character endpoints.
type CharacterHandler struct {
	svc *service.UserService
}

func NewCharacterHandler(svc *service.UserService) *CharacterHandler {
	return &CharacterHandler{svc: svc}
}

func (h *CharacterHandler) ListCharacters(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "user id not provided", http.StatusBadRequest)
		return
	}
	userID := parts[1]
	chars := h.svc.ListCharacters(userID)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chars)
}
