package promo

import (
	"testing"
)

func TestCalculatePromo(t *testing.T) {
	want := 200
	orderItems := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
	}
	got := CalculatePromo(orderItems)
	if got != want {
		t.Errorf("wanted %d but got %d", got, want)
	}

}
