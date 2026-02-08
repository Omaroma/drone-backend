package main

import (
	"log"
	"net/http"

	"github.com/Omaroma/drone-backend/controllers"
	"github.com/Omaroma/drone-backend/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/token", controllers.Token)
	mux.HandleFunc("/orders", middleware.Auth("enduser", controllers.SubmitOrder))
	mux.HandleFunc("/drone/reserve", middleware.Auth("drone", controllers.ReserveJob))
	mux.HandleFunc("/admin/orders", middleware.Auth("admin", controllers.ListOrders))
	mux.HandleFunc("/admin/orders/update", middleware.Auth("admin", controllers.UpdateOrder))
	mux.HandleFunc("/admin/drones", middleware.Auth("admin", controllers.ListDrones))
	mux.HandleFunc("/admin/drones/status", middleware.Auth("admin", controllers.UpdateDroneStatus))

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("error stating server", err)
	}
}
