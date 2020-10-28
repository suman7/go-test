package main

import (
	"app/promo"
	"fmt"
)

func main() {
	orderItems := map[string]int{
		"A": 1,
		"B": 1,
		"C": 1,
	}
	total := promo.CalculatePromo(orderItems)
	fmt.Println(total, orderItems)
}
