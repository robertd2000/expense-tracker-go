package repository

import "testing"

func TestNewRepository(t *testing.T) {
	repo := NewRepository("test.json")
	if repo == nil {		
		t.Errorf("got nil")
	}
}