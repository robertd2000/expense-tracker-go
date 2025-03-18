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

func TestAdd(t *testing.T) {
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

		addMultipleExpenses(repo, 10)

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

}

func TestDelete(t *testing.T) {
	t.Run("delete last", func(t *testing.T) {
		utils.Delete("test.json")
		repo := NewRepository("test.json")
		addMultipleExpenses(repo, 10)

		deleted, err := repo.Delete(10)

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		expenses, err := repo.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		if len(expenses) != 9 {
			t.Errorf("got total tasks %v, want %v", len(expenses), 9)
		}

		if id, _ := repo.GetLastID(); id != 9 {
			t.Errorf("got las ID %v, want %v", id, 9)
		}

		if expenses[0].ID != 1 {
			t.Errorf("got first ID %v, want %v", expenses[0].ID, 1)
		}

		if deleted.ID != 10 {
			t.Errorf("got deleted ID %v, want %v", deleted.ID, 10)
		}
	})

	t.Run("delete first", func(t *testing.T) {
		utils.Delete("test.json")
		repo := NewRepository("test.json")
		addMultipleExpenses(repo, 10)

		deleted, err := repo.Delete(1)

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		expenses, err := repo.GetAll()

		if err != nil {
			t.Errorf(err.Error())
			t.Errorf("got nil")
		}

		if len(expenses) != 9 {
			t.Errorf("got %v, want %v", len(expenses), 9)
		}

		if id, _ := repo.GetLastID(); id != 10 {
			t.Errorf("got las ID %v, want %v", id, 10)
		}

		if expenses[0].ID != 2 {
			t.Errorf("got first ID %v, want %v", expenses[0].ID, 2)
		}

		if deleted.ID != 1 {
			t.Errorf("got deleted ID %v, want %v", deleted.ID, 1)
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

func addMultipleExpenses(repository Repository, n int) {
	for i := 1; i <= n; i++ {
		repository.Save(*models.NewExpense("test"+fmt.Sprint(i), float64(i*100)))
	}
}