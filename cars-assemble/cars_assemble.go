package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (float64(successRate) / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// Calculate the number of cars produced per minute
	carsProducedPerMinute := productionRate / 60

	// Calculate the number of successful cars produced per minute
	successfulCarsPerMinute := int(float64(carsProducedPerMinute) * (successRate / 100))

	return successfulCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// Define the cost for producing 10 cars together and the cost for producing a single car
	costForGroupOf10 := 95000 // Cost for producing 10 cars together
	costPerCar := 10000       // Cost for producing a single car individually

	// Calculate the number of groups of 10 cars and the number of individual cars
	groupsOf10 := carsCount / 10
	individualCars := carsCount % 10

	// Calculate the total cost based on the number of groups and individual cars
	totalCost := uint(groupsOf10*costForGroupOf10 + individualCars*costPerCar)

	return totalCost
}
