package promo

import (
	"app/promoProp"
	"app/skuProp"
)

// Calculate total for promo of nItems
func PromoNItems(promo promoProp.Promotion, orderItems *map[string]int, total *int) {
	// Check if current promo is valid, and apply if valid
	sku := promo.SKUs[0]                                         // this kind of promo always has one SKU, so safe to take 0 index
	if val, ok := (*orderItems)[sku]; ok && val >= promo.Count { // As orderItems is map, lookup function is O(1)
		// If SKU found and order count is valid for this promo, lets update thet with promo offer
		promoApplyTimes := val / promo.Count // as same promo can be applied multiple times, if SKU has enough orders
		*total += promo.Value * promoApplyTimes
		// We wil also update the new order count, by subtracting SKU counts used the this promo
		if tmpCount := val - promo.Count*promoApplyTimes; tmpCount > 0 {
			(*orderItems)[sku] = tmpCount
		} else {
			// If all order count of SKU is used, we can safely remove th SKU from order, it will help to lookup faster
			delete((*orderItems), sku)
		}
	}

}

// Calculate total for promo of combineSKUs
func PromoCombineSKUs(promo promoProp.Promotion, orderItems *map[string]int, total *int) {
	// Check if this promo is valid, by checking if one of the promo SKU exists in order
	if val, ok := (*orderItems)[promo.SKUs[0]]; ok && val >= 1 {
		promoApplyTimes := val // as same promo can be applied multiple times, if SKU has enough orders
		// some Flag to apply promo, as we can break out of this loop if invalid, but not the function
		applyPromo := true
		// Loop through rest of the promo combination to check if they exist and valid
		for _, sku := range promo.SKUs[1:] {
			if val, ok = (*orderItems)[sku]; ok && val >= 1 {
				// We will consider the min of all SKU values, as the promo is a combination promo
				if promoApplyTimes < val {
					promoApplyTimes = val
				}
			} else {
				applyPromo = false
				break
			}
		}

		// If all SKUs exist in the order, we will apply this promo
		if applyPromo {
			*total += promo.Value * promoApplyTimes
			for _, sku := range promo.SKUs {
				if tmpCount := (*orderItems)[sku] - promoApplyTimes; tmpCount > 0 {
					(*orderItems)[sku] = tmpCount
				} else {
					delete((*orderItems), sku)
				}
			}
		}
	}
}

// Calculate total when no promo offer to apply
func NoPromoOrder(orderItems *map[string]int, total *int) {
	// Apply normal price to order
	for sku, count := range *orderItems {
		*total += skuProp.GetSKUPrice(sku) * count
	}
}

func CalculatePromo(orderItems map[string]int) int {
	total := 0

	// Apply promotion offer
	for _, promo := range promoProp.ActivePromotions {
		switch promo.Type {
		case promoProp.NItems:
			PromoNItems(promo, &orderItems, &total)
		case promoProp.CombineSKUs:
			PromoCombineSKUs(promo, &orderItems, &total)
		}
	}

	// Update total when no promotion to apply
	NoPromoOrder(&orderItems, &total)
	return total
}
