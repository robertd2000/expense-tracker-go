package repository

import (
	"testing"

	"github.com/robertd2000/expense-tracker/interval/utils"
)

func TestNewRepository(t *testing.T) {
	utils.Delete("test.json")
	
	repo := NewRepository("test.json")
	if repo == nil {		
		t.Errorf("got nil")
	}
}