package service

import (
	"testing"
	"time"

	"fmt"

	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/repository"
	"github.com/robertd2000/expense-tracker/interval/utils"
)

func TestAddOne(t *testing.T) {
	utils.Delete("test.json")

	expenseRepository := repository.NewRepository("test.json")
	expenseService := NewExpenseService(expenseRepository)

	got, err := expenseService.Add("test", 1.0)

	if err != nil {
		t.Errorf(err.Error())
		t.Errorf("got nil")
	}
	want := &models.Expense{
		ID:      1,
		Details: "test",
		Amount:  1.0,
		Date:    time.Now(),
	}

	if got.Amount != want.Amount || got.Details != want.Details || got.ID != want.ID {
		t.Errorf("got %v, want %v", got, want)
	}
}

func MockExpenseTasks() []models.Expense {
	tasks := make([]models.Expense, 0, 10) 

    // Заполняем срез задачами
    for i := 1; i <= 10; i++ {
        tasks = append(tasks, models.Expense{
            ID:     i,       
            Amount: float64(i * 100),
			Details: "test" + fmt.Sprint(i),
        })
    }

	return tasks
}

func TestAddMultiple(t *testing.T) {
	utils.Delete("test.json")

	expenseRepository := repository.NewRepository("test.json")
	expenseService := NewExpenseService(expenseRepository)

	for i := 1; i <= 10; i++ {
		expenseService.Add("test" + fmt.Sprint(i), float64(i * 100))
	}

	expenses, err := expenseRepository.GetAll()

	if err != nil {
		t.Errorf(err.Error())
		t.Errorf("got nil")
	}

	if len(expenses) != 10 {
		t.Errorf("got %v, want %v", len(expenses), 10)
	}

	mockTasks := MockExpenseTasks()
	for i := 0; i < 10; i++ {
		if expenses[i].Amount != mockTasks[i].Amount || expenses[i].Details != mockTasks[i].Details || expenses[i].ID != mockTasks[i].ID {
			t.Errorf("got %v, want %v", expenses[i], mockTasks[i])
		}
	}

}