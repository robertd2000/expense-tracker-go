package service

import (
	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/repository"
)

type ExpenseService interface {
	Add(details string, amount float64) (*models.Expense, error)
	GetAll() ([]models.Expense, error)
	Delete(id int) (*models.Expense, error)
}

type expenseService struct {
	repository repository.Repository
}


func NewExpenseService(repository repository.Repository) ExpenseService {
	return &expenseService{
		repository: repository}
}

func (e *expenseService) Add(details string, amount float64) (*models.Expense, error) {
	expense := models.NewExpense(details, amount)

	expense, err := e.repository.Save(*expense)

	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (e *expenseService) GetAll() ([]models.Expense, error) {
	return e.repository.GetAll()
}

func (e *expenseService) Delete(id int) (*models.Expense, error) {
	return e.repository.Delete(id)
}

