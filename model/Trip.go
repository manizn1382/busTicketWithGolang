package model

import "time"

type Trip struct {
	TripId        int
	DepartureTime time.Time
	ArrivalTime time.Time
	Price float32
	Status string
	Origin string
	Dest string
	Distance float32
}