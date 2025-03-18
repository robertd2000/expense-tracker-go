package repository

import (
	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/utils"
)

type Repository interface {
	Save(expense models.Expense) (models.Expense, error)
	GetAll() ([]models.Expense, error)
}

type repository struct {
	sourceFile string
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Save(expense models.Expense) (models.Expense, error) {
	return expense, nil
}

func (r *repository) getData() (*models.ExpenseDB, error) {
	stream, err := utils.ReadFromJSON(r.sourceFile)

	if err != nil {
		return nil, err
	}

	expenseData, err := utils.DeserializeFromJSON[models.ExpenseDB](stream)

	if err != nil {
		return nil, err
	}

	return &expenseData, nil
}

func (r *repository) GetAll() ([]models.Expense, error) {
	expenseData, err := r.getData()

	if err != nil {
		return nil, err
	}

	return expenseData.Expenses, nil
}

func (r *repository) GetLastID() (int, error) {
	expenseData, err := r.getData()

	if err != nil {
		return -1, err
	}

	return expenseData.LastID, nil
}