package utils

import "testing"

func CheckParams(t testing.TB, delId, expId, length, expLength, firstId, expFirstId, lastId, expLastId int) {
	t.Helper()
	if delId != expId {
		t.Errorf("got deleted ID %v, want %v", delId, expId)
	}

	if length != expLength {
		t.Errorf("got length %v, want %v", length, expLength)
	}

	if firstId != expFirstId {
		t.Errorf("got first ID %v, want %v", firstId, expFirstId)
	}

	if lastId != expLastId {
		t.Errorf("got last ID %v, want %v", lastId, expLastId)
	}
}