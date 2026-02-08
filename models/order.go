package models

import "time"

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Reserved  OrderStatus = "reserved"
	PickedUp  OrderStatus = "picked_up"
	Delivered OrderStatus = "delivered"
	Failed    OrderStatus = "failed"
)

type Order struct {
	ID          string      `json:"id"`
	Owner       string      `json:"owner"`
	Origin      Location    `json:"origin"`
	Destination Location    `json:"destination"`
	Status      OrderStatus `json:"status"`
	AssignedTo  string      `json:"assigned_to,omitempty"`
	CreatedAt   time.Time   `json:"created_at"`
}
