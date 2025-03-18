package service

import (
	"testing"
	"time"

	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/repository"
)

func TestAdd(t *testing.T) {

	expenseRepository := repository.NewRepository("test.json")
	expenseService := NewExpenseService(expenseRepository)

	got, err := expenseService.Add("test", 1.0)

	if err != nil {
		t.Errorf(err.Error())
		t.Errorf("got nil")
	}
	want := &models.Expense{
		ID:      2,
		Details: "test",
		Amount:  1.0,
		Date:    time.Now(),
	}

	if got.Amount != want.Amount || got.Details != want.Details || got.ID != want.ID {
		t.Errorf("got %v, want %v", got, want)
	}
}