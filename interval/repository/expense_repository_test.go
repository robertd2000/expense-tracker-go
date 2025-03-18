package repository

import (
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
	checkData := func(t testing.TB, got, want interface{}) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("add one task", func(t *testing.T) {
		repo := NewRepository("test.json")
		repo.Save(*models.NewExpense("test", 1.0))
		expenses, err := repo.GetAll()

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

		checkData(t, expenses[0], want)
	})
}