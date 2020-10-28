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
func PromoNItems(promo Promotion, orderItems *map[string]int, total *int) {
	// Check if current promo is valid, and apply if valid
	sku := promo.SKUs[0]                                         // this kind of promo always has one SKU, so safe to take 0 index
	if val, ok := (*orderItems)[sku]; ok && val >= promo.Count { // As orderItems is map, lookup function is O(1)
		// If SKU found and order count is valid for this promo, lets update thet with promo offer
		*total += promo.Value
		// We wil also update the new order count, by subtracting SKU counts used the this promo
		if tmpCount := val - promo.Count; tmpCount > 0 {
			(*orderItems)[sku] = tmpCount
		} else {
			// If all order count of SKU is used, we can safely remove th SKU from order, it will help to lookup faster
			delete((*orderItems), sku)
		}
	}

}

// Calculate total for promo of combineSKUs
func PromoCombineSKUs(promo Promotion, orderItems *map[string]int, total *int) {
	fmt.Println(*total)
}

func CalculatePromo(orderItems map[string]int) int {
	total := 0

	// Apply promotion offer
	for _, promo := range activePromotions {
		switch promo.Type {
		case nItems:
			PromoNItems(promo, &orderItems, &total)
		case combineSKUs:
			PromoCombineSKUs(promo, &orderItems, &total)
		}
	}

	return total
}
