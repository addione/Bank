package models

import "time"

type Transaction struct {
	Amount     float64
	SenderId   int64
	ReceiverId int64
	Time       time.Time
	Status     string
}
