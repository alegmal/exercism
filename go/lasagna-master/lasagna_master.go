package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * averagePreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
    var sauce float64
    for _, layer := range layers {
        switch layer {
        case "noodles":
            noodles += 50
        case "sauce":
            sauce += 0.2
        }
    }
    return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))
	for i, amount := range amounts {
		scaledAmounts[i] = amount * float64(portions) / 2.0
	}

	return scaledAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
