package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Omaroma/drone-backend/services"
)

func ReserveJob(w http.ResponseWriter, r *http.Request) {
	order, err := services.ReserveOrder("drone-1")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(order)
}
