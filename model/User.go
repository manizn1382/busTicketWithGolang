package model

import "time"

type User struct {
	UserId   int
	Name     string
	Role     string
	PassHash string
	Phone    string
	CreateAt time.Time
	NationalId string
}