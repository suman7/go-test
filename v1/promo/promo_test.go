package promo

import (
	"testing"
)

func TestCalculatePromo(t *testing.T) {
	want := 100
	got := 200
	if got != want {
		t.Errorf("wanted %d but got %d", got, want)
	}

}
