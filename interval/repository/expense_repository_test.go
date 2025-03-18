package repository

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/robertd2000/expense-tracker/interval/models"
	"github.com/robertd2000/expense-tracker/interval/utils"
)

func TestNewRepository(t *testing.T) {
	utils.Delete("test.json")
	
	got := NewRepository("test.json")
	if got == nil {		
		t.Errorf("got nil")
	}

	want := &repository{sourceFile: "test.json", tasks: nil, lastID: 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCRUD(t *testing.T) {
	checkData := func(t testing.TB, got, want models.Expense) {
		t.Helper()
		if got.Amount != want.Amount || got.Details != want.Details || got.ID != want.ID {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("add one task", func(t *testing.T) {
		utils.Delete("test.json")
		repo := NewRepository("test.json")
		repo.Save(*models.NewExpense("test", 1.0))
		expenses, err := repo.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := models.Expense{
			ID:      1,
			Details: "test",
			Amount:  1.0,
			Date:    time.Now(),
		}

		checkData(t, expenses[0], want)
	})

	t.Run("add 10 tasks", func(t *testing.T) {
		utils.Delete("test.json")
		repo := NewRepository("test.json")

		for i := 1; i <= 10; i++ {
			repo.Save(*models.NewExpense("test"+fmt.Sprint(i), float64(i*100)))
		}

		expenses, err := repo.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		want := MockExpenseTasks()

		if len(expenses) != 10 {
			t.Errorf("got %v, want %v", len(expenses), 10)
		}

		for i := range 10 {
			checkData(t, expenses[i], want[i])
		}
	})
	t.Run("delete last", func(t *testing.T) {
		repo := NewRepository("test.json")

		repo.Delete(10)

		expenses, err := repo.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		if len(expenses) != 9 {
			t.Errorf("got %v, want %v", len(expenses), 9)
		}

		if id, _ := repo.GetLastID(); id != 9 {
			t.Errorf("got %v, want %v", id, 9)
		}
	})
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