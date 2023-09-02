package pkg

import (
	"testing"
)

func TestAddLength(t *testing.T) {
	got := addLength([]rune("・"), "・")
	want := "・・"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
