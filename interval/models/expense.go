package models

import "time"

type Expense struct {
	ID      int     `json:"id"`
	Details string  `json:"details"`
	Amount  float64 `json:"amount"`
	Date    time.Time  `json:"date"`
}

func NewExpense(details string, amount float64) *Expense {
	return &Expense{
		Details: details,
		Amount:  amount,
		Date:    time.Now(),
	}
}