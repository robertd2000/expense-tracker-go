package repository

import (
	"fmt"

	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/utils"
)

type Repository interface {
	Save(expense models.Expense) (*models.Expense, error)
	GetAll() ([]models.Expense, error)
	GetLastID() (int, error)
}

type repository struct {
	sourceFile string
	tasks      []models.Expense
	lastID     int
}

func NewRepository(sourceFile string) Repository {
	repo := &repository{
		sourceFile: sourceFile,
	}

	repo.Init()

	return repo
}

func (r *repository) Init() {
	data, err := r.getData()

	if err != nil {
		return
	}

	r.tasks = data.Expenses
	r.lastID = data.LastID

}

func (r *repository) Save(expense models.Expense) (*models.Expense, error) {
	id, err := r.GetLastID()

	if err != nil {
		return nil, err
	}

	expense.ID = id + 1

	r.tasks = append(r.tasks, expense)
	fmt.Println(r.tasks)

	r.lastID = expense.ID

	if err := r.commit(); err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *repository) getData() (*models.ExpenseDB, error) {
	stream, err := utils.ReadFromJSON(r.sourceFile)

	if err != nil {
		return &models.ExpenseDB{}, nil
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

func (r *repository) commit() error {
	db := models.ExpenseDB{
		Expenses: r.tasks,
		LastID:   r.lastID,
	} 
	s, err := utils.SerializeToJSON(db)
	if err != nil {
		return fmt.Errorf("unable to serialize task: %w", err)
	}

	if err := utils.SaveToJSON(r.sourceFile, s); err != nil {
		return fmt.Errorf("unable to save to JSON: %w", err)
	}

	return nil
}