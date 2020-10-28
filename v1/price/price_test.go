package price

import (
	"testing"
)

func TestGetSKUPrice(t *testing.T) {

	t.Run("Price A", func(t *testing.T) {
		want := 50
		got := GetSKUPrice(A)
		if got != want {
			t.Errorf("wanted %d but got %d", want, got)
		}
	})

}
