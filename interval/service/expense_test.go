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

func TestAddMultiple(t *testing.T) {
	utils.Delete("test.json")

	expenseRepository := repository.NewRepository("test.json")
	expenseService := NewExpenseService(expenseRepository)

	addMultipleExpenses(expenseService, 10)

	expenses, err := expenseService.GetAll()

	if err != nil {
		t.Errorf(err.Error())
		t.Errorf("got nil")
	}

	if len(expenses) != 10 {
		t.Errorf("got %v, want %v", len(expenses), 10)
	}

	mockTasks := MockExpenseTasks()
	for i := range 10 {
		if expenses[i].Amount != mockTasks[i].Amount || expenses[i].Details != mockTasks[i].Details || expenses[i].ID != mockTasks[i].ID {
			t.Errorf("got %v, want %v", expenses[i], mockTasks[i])
		}
	}

	lastId, _ := expenseRepository.GetLastID()

	if lastId != 10 {
		t.Errorf("got %v, want %v", lastId, 10)
	}
}

func TestDelete(t *testing.T) {
	t.Run("delete last", func(t *testing.T) {
		utils.Delete("test.json")

		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		deleted, err := expenseService.Delete(10)

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		got, err := expenseService.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		if len(got) != 9 {
			t.Errorf("got total tasks %v, want %v", len(got), 9)
		}

		if got[0].ID != 1 {
			t.Errorf("got first ID %v, want %v", got[0].ID, 1)
		}

		if deleted.ID != 10 {
			t.Errorf("got deleted ID %v, want %v", deleted.ID, 10)
		}

		want := MockExpenseTasks()[:9]

		for i := range 9 {
			if got[i].Amount != want[i].Amount || got[i].Details != want[i].Details || got[i].ID != want[i].ID {
				t.Errorf("got %v, want %v", got[i], want[i])
			}
		}
	})
}

func addMultipleExpenses(expenseService ExpenseService, n int) {
	for i := 1; i <= n; i++ {
		expenseService.Add("test"+fmt.Sprint(i), float64(i*100))
	}
}


func MockExpenseTasks() []models.Expense {
	tasks := make([]models.Expense, 0, 10) 

    for i := 1; i <= 10; i++ {
        tasks = append(tasks, models.Expense{
            ID:     i,       
            Amount: float64(i * 100),
			Details: "test" + fmt.Sprint(i),
        })
    }

	return tasks
}