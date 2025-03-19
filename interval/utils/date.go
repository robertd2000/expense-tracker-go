package utils

import "time"

func GetCurrentMonth() int {
	now := time.Now()

	currentMonth := now.Month()

	currentMonthNumber := int(currentMonth)

	return currentMonthNumber
}