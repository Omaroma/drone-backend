package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Omaroma/drone-backend/services"
)

func Token(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	token, _ := services.GenerateToken(req.Name, req.Role)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
