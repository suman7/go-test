package promo

import (
	"testing"
)

func TestCalculatePromo(t *testing.T) {

	assertSuccess := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("Scenario A", func(t *testing.T) {
		want := 100
		orderItems := map[string]int{
			"A": 1,
			"B": 1,
			"C": 1,
		}
		got := CalculatePromo(orderItems)
		assertSuccess(t, got, want)
	})

	t.Run("Scenario B", func(t *testing.T) {
		want := 370
		orderItems := map[string]int{
			"A": 5,
			"B": 5,
			"C": 1,
		}
		got := CalculatePromo(orderItems)
		assertSuccess(t, got, want)
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
		assertSuccess(t, got, want)
	})

}
