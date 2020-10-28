package promo

import (
	"fmt"
)

// Promotion structure
type Promotion struct {
	SKUs  []string // e.g. ["A"] or ["C", "D"]
	Count int      // SKU's quantity in the promotional offer, e.g 3 , when 3 of SKU:A bought together
	Value int      // the offer value, e.g. 130
	Type  string   // the promotion type, e.g. nItems: when n items bought together
}

// Enums for easy handling and avoiding type
const (
	// Promotion types
	nItems      string = "nItems"
	combineSKUs string = "combineSKUs"

	// SKU types
	A string = "A"
	B string = "B"
	C string = "C"
	D string = "D"
)

// Active promotions, can be managed by admin ui in real scenario
var activePromotions = map[string]Promotion{
	"nItemsA":   Promotion{[]string{A}, 3, 130, nItems},        // 3 of A's for 130
	"nItemsB":   Promotion{[]string{B}, 2, 45, nItems},         // 2 of B's for 45
	"combineCD": Promotion{[]string{C, D}, 0, 30, combineSKUs}, // C & D for 30
}

// Calculate total for promo of nItems
func PromoNItems() {
	fmt.Println("PromoNItems called")
}

// Calculate total for promo of combineSKUs
func PromoCombineSKUs() {
	fmt.Println("PromoCombineSKUs called")
}

func CalculatePromo(orderItems map[string]int) int {
	total := 0

	for _, promo := range activePromotions {
		switch promo.Type {
		case nItems:
			PromoNItems()
		case combineSKUs:
			PromoCombineSKUs()
		}
	}

	return total
}
