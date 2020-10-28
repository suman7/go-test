package promo

import (
	"testing"
)

func TestCalculatePromo(t *testing.T) {

	t.Run("Scenario A", func(t *testing.T) {
		want := 200
		orderItems := map[string]int{
			"A": 1,
			"B": 1,
			"C": 1,
		}
		got := CalculatePromo(orderItems)
		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

	t.Run("Scenario B", func(t *testing.T) {
		want := 370
		orderItems := map[string]int{
			"A": 5,
			"B": 5,
			"C": 1,
		}
		got := CalculatePromo(orderItems)
		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

	t.Run("Scenario C", func(t *testing.T) {
		want := 280
		orderItems := map[string]int{
			"A": 3,
			"B": 5,
			"C": 1,
			"D": 1,
		}
		got := CalculatePromo(orderItems)
		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

}
