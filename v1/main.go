package main

import (
	"app/promo"
)

func main() {
	orderItems := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
	}
	promo.CalculatePromo(orderItems)
}
