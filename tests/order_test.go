package tests

import (
	"testing"

	"github.com/Omaroma/drone-backend/models"
	"github.com/Omaroma/drone-backend/services"
)

func TestCreateOrder(t *testing.T) {
	o := models.Order{}
	order := services.CreateOrder("alice", o)

	if order.Owner != "alice" {
		t.Fatal("owner not set")
	}
}
