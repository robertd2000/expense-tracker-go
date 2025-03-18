package repository

import (
	"reflect"
	"testing"

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