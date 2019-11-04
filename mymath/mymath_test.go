package mymath

import "testing"

func TestAddInt(t *testing.T) {
	want := 3
	if got := AddInt(1, 2); got != want {
		t.Errorf("Add(1, 2) = %q, want %q", got, want)
	}
}
