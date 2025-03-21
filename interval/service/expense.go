package service

import (
	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/repository"
)

type ExpenseService interface {
	GetAll() ([]models.Expense, error)
	Add(details string, amount float64) (*models.Expense, error)
	Delete(id int) (*models.Expense, error)
	Update(id int, expense models.Expense) (*models.Expense, error)
	GetSummary(filterDate ...int) (float64, error)
}

type expenseService struct {
	repository repository.Repository
}

func NewExpenseService(repository repository.Repository) ExpenseService {
	return &expenseService{
		repository: repository}
}

func (e *expenseService) GetAll() ([]models.Expense, error) {
	return e.repository.GetAll()
}

func (e *expenseService) Add(details string, amount float64) (*models.Expense, error) {
	expense := models.NewExpense(details, amount)

	expense, err := e.repository.Save(*expense)

	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (e *expenseService) Delete(id int) (*models.Expense, error) {
	return e.repository.Delete(id)
}

func (e *expenseService) Update(id int, expense models.Expense) (*models.Expense, error) {
	return e.repository.Update(id, expense)
}

func (e *expenseService) GetSummary(filterDate ...int) (float64, error) {
	return e.repository.GetSummary(filterDate...)
}
