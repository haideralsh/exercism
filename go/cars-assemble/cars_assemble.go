package cars


// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionRatePerMin := float64(productionRate) / 60
	successRatePerMin := successRate

	return int((productionRatePerMin * successRatePerMin) / 100.00)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	remainder := carsCount % 10
	
	return uint((tens * 95000) + (remainder * 10000))
}
