package model

import "time"

type Payment struct {
	PaymentId int
	TicketId  int
	Amount    float32
	PayType   string
	PayStatus string
	CreateAt  time.Time
}