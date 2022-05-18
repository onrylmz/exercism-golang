package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	message := " is clearly the better choice."

	var recommendedVehicle string
	if option1 < option2 {
		recommendedVehicle = option1
	} else {
		recommendedVehicle = option2
	}

	return recommendedVehicle + message
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var i float64
	if age < 3 {
		i = .8
	} else if age < 10 {
		i = .7
	} else {
		i = .5
	}

	return i * originalPrice
}
