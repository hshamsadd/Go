package handler

import (
	"Go/my-go-app/internal/service"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Svc *service.UserService
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	
	user, err := h.Svc.GetProfile(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}