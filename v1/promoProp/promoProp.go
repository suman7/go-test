package promoProp

import (
	"app/skuProp"
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
	NItems      string = "nItems"
	CombineSKUs string = "combineSKUs"
)

// Active promotions, can be managed by admin ui in real scenario
var ActivePromotions = map[string]Promotion{
	"nItemsA":   Promotion{[]string{skuProp.A}, 3, 130, NItems},                // 3 of A's for 130
	"nItemsB":   Promotion{[]string{skuProp.B}, 2, 45, NItems},                 // 2 of B's for 45
	"combineCD": Promotion{[]string{skuProp.C, skuProp.D}, 0, 30, CombineSKUs}, // C & D for 30
}
