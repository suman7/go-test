package promo

import (
	"testing"
)

func TestCalculatePromo(t *testing.T) {
	want := 200
	got := CalculatePromo()
	if got != want {
		t.Errorf("wanted %d but got %d", got, want)
	}

}
