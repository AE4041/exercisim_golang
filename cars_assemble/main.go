package carsassemble

const (
	costPerTenCars   = 95000
	costPerRemaining = 10000
	carsPerGroup     = 10
)

// CalculateWorkingCarsPerHour calculates how many working cars are produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate) / 100) * successRate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost calculates the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount < 0 {
		return 0
	}
	tens := uint(carsCount) / carsPerGroup
	reminder := uint(carsCount) % carsPerGroup
	return (tens * costPerTenCars) + (reminder * costPerRemaining)
}
