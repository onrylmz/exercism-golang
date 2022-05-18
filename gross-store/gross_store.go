package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	var units = make(map[string]int)

	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, unitExists := units[unit]

	if !unitExists {
		return false
	}

	_, itemExists := bill[item]

	if itemExists {
		bill[item] += value
	} else {
		bill[item] = value
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	billItem, billExists := bill[item]

	if !billExists {
		return false
	}

	unitItem, unitExists := units[unit]
	if !unitExists {
		return false
	}

	if billItem < unitItem {
		return false
	} else if billItem == unitItem {
		delete(bill, item)
		return true
	} else {
		bill[item] -= unitItem
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	billItem, exists := bill[item]

	return billItem, exists
}
