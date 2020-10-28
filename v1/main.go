package main

import (
	"app/promo"
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	orderItems := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
	}
	promo.CalculatePromo(orderItems)
}
