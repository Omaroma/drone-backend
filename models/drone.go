package models

import "time"

type Drone struct {
	ID        string    `json:"id"`
	Location  Location  `json:"location"`
	Broken    bool      `json:"broken"`
	UpdatedAt time.Time `json:"updated_at"`
}
