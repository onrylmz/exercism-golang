package cars

const CAR_PER_HOUR float64 = 221.0

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	} else if 1 <= speed && speed <= 4 {
		return 1.0
	} else if 5 <= speed && speed <= 8 {
		return 0.9
	} else {
		return 0.77
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return SuccessRate(speed) * float64(speed) * CAR_PER_HOUR
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	productionRatePerHour := CalculateProductionRatePerHour(speed)

	return int(productionRatePerHour / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	productionRatePerHour := CalculateProductionRatePerHour(speed)

	if productionRatePerHour > limit {
		return limit
	}

	return productionRatePerHour
}
