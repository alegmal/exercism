package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}

	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	const text = " is clearly the better choice."

	if option1 < option2 {
		return option1 + text
	}

	return option2 + text
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var reductionPercent float64

	if age < 3 {
		reductionPercent = 0.8
	} else if age >= 3 && age < 10 {
		reductionPercent = 0.7
	} else if age >= 10 {
		reductionPercent = 0.5
	}

	return originalPrice * reductionPercent
}
