package service

import (
	"testing"

	"github.com/robertd2000/expense-tracker/interval/models"
)

func TestAdd(t *testing.T) {
	expenseService := NewExpenseService()

	got := expenseService.Add("test", 1.0)
	want := models.Expense{}

	if got != want {
		t.Errorf("got %w want %w", got, want)
	}
}