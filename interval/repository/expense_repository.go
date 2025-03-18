package repository

import "github.com/robertd2000/expense-tracker/interval/models"

type Repository interface {
	Save(expense models.Expense) (models.Expense, error)
	GetAll() ([]models.Expense, error)
}

type repository struct {}

func NewRepository() Repository {
	return &repository{}
}

