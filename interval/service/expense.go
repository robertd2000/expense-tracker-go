package service

import (
	"github.com/robertd2000/expense-tracker/interval/models"
)

type ExpenseService interface {
	Add(details string, amount float64) (models.Expense, error)
}

type expenseService struct{}

func NewExpenseService() ExpenseService {
	return &expenseService{}
}

func (e *expenseService) Add(details string, amount float64) (models.Expense, error) {
	return *models.NewExpense(1, details, amount), nil
}