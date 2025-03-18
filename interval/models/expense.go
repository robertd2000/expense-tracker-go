package models

import "time"

type Expense struct {
	ID      int     `json:"id"`
	Details string  `json:"details"`
	Amount  float64 `json:"amount"`
	Date    time.Time  `json:"date"`
}