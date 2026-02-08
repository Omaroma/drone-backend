package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Omaroma/drone-backend/models"
	"github.com/Omaroma/drone-backend/services"
)

func SubmitOrder(w http.ResponseWriter, r *http.Request) {
	var req models.Order
	json.NewDecoder(r.Body).Decode(&req)

	order := services.CreateOrder("user", req)
	json.NewEncoder(w).Encode(order)
}
