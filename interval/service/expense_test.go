package service

import (
	"testing"

	"github.com/robertd2000/expense-tracker/interval/models"
)

func TestAdd(t *testing.T) {
	expenseService := NewExpenseService()

	got, err := expenseService.Add("test", 1.0)

	if err != nil {
		t.Errorf("got nil")
	}
	want := models.Expense{}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}