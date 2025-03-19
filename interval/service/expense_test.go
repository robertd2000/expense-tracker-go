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

	expenses := getExpenses(expenseService, t)

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
	checkData := func(t testing.TB, got, want models.Expense) {
		t.Helper()
		if got.Amount != want.Amount || got.Details != want.Details || got.ID != want.ID {
			t.Errorf("got %v, want %v", got, want)
		}
	}

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

		got := getExpenses(expenseService, t)
		lastId, _ := expenseRepository.GetLastID()

		utils.CheckParams(t, deleted.ID, 10, len(got), 9, got[0].ID, 1, lastId, 9)

		want := MockExpenseTasks()[:9]

		for i := range 9 {
			checkData(t, got[i], want[i])
		}
	})

	t.Run("delete first", func(t *testing.T) {
		utils.Delete("test.json")

		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		deleted, err := expenseService.Delete(1)

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		got := getExpenses(expenseService, t)
		lastId, _ := expenseRepository.GetLastID()

		utils.CheckParams(t, deleted.ID, 1, len(got), 9, got[0].ID, 2, lastId, 10)

		want := MockExpenseTasks()[1:10]

		for i := range 9 {
			checkData(t, got[i], want[i])
		}
	})
}

func TestUpdate(t *testing.T) {
	checkData := func(t testing.TB, got, want models.Expense) {
		t.Helper()
		if got.Amount != want.Amount || got.Details != want.Details || got.ID != want.ID {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("update one task details", func(t *testing.T) {
		utils.Delete("test.json")

		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		expenseService.Update(1, models.Expense{Details: "updated"})
		expenses, err := expenseService.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := models.Expense{
			ID:      1,
			Details: "updated",
			Amount:  100,
			Date:    time.Now(),
		}

		checkData(t, expenses[0], want)
	})
	
	t.Run("update multiple tasks details", func(t *testing.T) {
		utils.Delete("test.json")
		
		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		for i := 3; i <= 6; i++ {
			expenseService.Update(i, models.Expense{Details: "updated" + fmt.Sprint(i)})
		}

		expenses, err := expenseService.GetAll()

		for i := 3; i <= 6; i++ {
			want := models.Expense{
				ID:      i,
				Details: "updated" + fmt.Sprint(i),
				Amount:  float64(i * 100),
				Date:    time.Now(),
			}

			checkData(t, expenses[i-1], want)
		}

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}
	})

	t.Run("update one task amount", func(t *testing.T) {
		utils.Delete("test.json")
		
		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		expenseService.Update(1, models.Expense{Amount: 111})
		expenses, err := expenseService.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := models.Expense{
			ID:      1,
			Details: "test1",
			Amount:  111,
			Date:    time.Now(),
		}

		checkData(t, expenses[0], want)
	})

	t.Run("update multiple tasks amount", func(t *testing.T) {
		utils.Delete("test.json")
		
		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		for i := 3; i <= 6; i++ {
			expenseService.Update(i, models.Expense{Amount: float64(i * 111)})
		}

		expenses, err := expenseService.GetAll()

		for i := 3; i <= 6; i++ {
			want := models.Expense{
				ID:      i,
				Details: "test" + fmt.Sprint(i),
				Amount:  float64(i * 111),
				Date:    time.Now(),
			}

			checkData(t, expenses[i-1], want)
		}

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}
	})
}

func TestSummary(t *testing.T) {
	checkData := func(t testing.TB, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("get summary of 10 tasks from 100 to 1000", func(t *testing.T) {
		utils.Delete("test.json")
		
		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)

		summary, err := expenseService.GetSummary()
		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := 5500.0

		checkData(t, summary, want)
	})
	
	t.Run("get summary of 9 tasks from 100 to 1000 with ID 5 deleted", func(t *testing.T) {
		utils.Delete("test.json")
			
		expenseRepository := repository.NewRepository("test.json")
		expenseService := NewExpenseService(expenseRepository)

		addMultipleExpenses(expenseService, 10)
		expenseService.Delete(5)

		summary, err := expenseService.GetSummary()
		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := 5000.0

		checkData(t, summary, want)
	})
}

func getExpenses(expenseService ExpenseService, t *testing.T)  []models.Expense {
	got, err := expenseService.GetAll()

	if err != nil {
		t.Errorf(err.Error())
		t.Errorf("got nil")
	}

	return got
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