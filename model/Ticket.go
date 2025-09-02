package model

import "time"

type Ticket struct {
	TicketId int
	TripId   int
	UserId   int
	SeatId   int
	BookTime time.Time
	Status string
}