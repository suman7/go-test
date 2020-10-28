package skuProp

// Enums for easy handling and avoiding type
const (
	// SKU types
	A string = "A"
	B string = "B"
	C string = "C"
	D string = "D"
)

// Should be fetched from DB, in real application
var skuPrice = map[string]int{
	A: 50,
	B: 30,
	C: 20,
	D: 15,
}

// Should be a DB fetch logic, in real application
func GetSKUPrice(sku string) int {
	return skuPrice[sku]
}
