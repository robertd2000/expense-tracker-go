package service

import "testing"

func TestAdd(t *testing.T) {
	expenseService := NewExpenseService()

	got := expenseService.Add("test", 1.0)
	want := Expense{}

	if got != want {
		t.Errorf("got %w want %w", got, want)
	}
}