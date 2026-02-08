package store

import (
	"sync"

	"github.com/Omaroma/drone-backend/models"
)

type MemoryStore struct {
	Mu     sync.Mutex
	Orders map[string]*models.Order
	Drones map[string]*models.Drone
}

var Store = &MemoryStore{
	Orders: map[string]*models.Order{},
	Drones: map[string]*models.Drone{},
}
