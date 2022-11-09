package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if !found(units, unit) {
		return false
	}

	bill[item] = bill[item] + units[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if !found(units, unit) || !found(bill, item) {
		return false
	}

	newVal := bill[item] - units[unit]

	switch {
	case newVal < 0:
		return false

	case newVal == 0:
		delete(bill, item)
		return true

	default:
		bill[item] = bill[item] - units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if !found(bill, item) {
		return 0, false
	}

	return bill[item], true
}

func found[V any](m map[string]V, key string) bool {
	_, ok := m[key]

	return ok
}
