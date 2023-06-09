package cars

const singleCarManufactoringCost = 10000
const tenCarManufactoringCost = 95000
const minutesInHour = 60

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / minutesInHour)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	if carsCount < 10 {
		return uint(singleCarManufactoringCost * carsCount)
	}

	groupOfTenCars := carsCount / 10
	remaningSingleCars := carsCount % 10

	return uint(groupOfTenCars*tenCarManufactoringCost + remaningSingleCars*singleCarManufactoringCost)
}
