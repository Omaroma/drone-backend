package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Omaroma/drone-backend/models"
	"github.com/Omaroma/drone-backend/store"
)

func ListOrders(w http.ResponseWriter, r *http.Request) {
	store.Store.Mu.Lock()
	defer store.Store.Mu.Unlock()

	orders := make([]*models.Order, 0, len(store.Store.Orders))
	for _, o := range store.Store.Orders {
		orders = append(orders, o)
	}

	json.NewEncoder(w).Encode(orders)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID          string           `json:"id"`
		Origin      *models.Location `json:"origin,omitempty"`
		Destination *models.Location `json:"destination,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	store.Store.Mu.Lock()
	defer store.Store.Mu.Unlock()

	order, ok := store.Store.Orders[req.ID]
	if !ok {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	if req.Origin != nil {
		order.Origin = *req.Origin
	}
	if req.Destination != nil {
		order.Destination = *req.Destination
	}

	json.NewEncoder(w).Encode(order)
}

func ListDrones(w http.ResponseWriter, r *http.Request) {
	store.Store.Mu.Lock()
	defer store.Store.Mu.Unlock()

	drones := make([]*models.Drone, 0, len(store.Store.Drones))
	for _, d := range store.Store.Drones {
		drones = append(drones, d)
	}

	json.NewEncoder(w).Encode(drones)
}

func UpdateDroneStatus(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     string `json:"id"`
		Broken bool   `json:"broken"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	store.Store.Mu.Lock()
	defer store.Store.Mu.Unlock()

	drone, ok := store.Store.Drones[req.ID]
	if !ok {
		http.Error(w, "drone not found", http.StatusNotFound)
		return
	}

	drone.Broken = req.Broken
	json.NewEncoder(w).Encode(drone)
}
