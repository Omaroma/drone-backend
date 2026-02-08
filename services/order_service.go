package services

import (
	"errors"
	"time"

	"github.com/Omaroma/drone-backend/models"
	"github.com/Omaroma/drone-backend/store"
)

func CreateOrder(owner string, o models.Order) *models.Order {
	o.ID = time.Now().Format("20060102150405")
	o.Owner = owner
	o.Status = models.Pending
	o.CreatedAt = time.Now()

	store.Store.Orders[o.ID] = &o
	return &o
}

func ReserveOrder(droneID string) (*models.Order, error) {
	for _, o := range store.Store.Orders {
		if o.Status == models.Pending {
			o.Status = models.Reserved
			o.AssignedTo = droneID
			return o, nil
		}
	}
	return nil, errors.New("no available orders")
}
