package lasagna


func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	saucesCount := 0.0
	noodlesCount := 0

	for _, layer := range layers {
		if layer == "noodles" { noodlesCount++ }
		if layer == "sauce" { saucesCount++ }
	}


	return noodlesCount * 50, saucesCount * 0.2
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friend []string, own []string) {
	last := friend[len(friend) - 1]

	own[len(own) - 1] = last
}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var adjustedToPortions []float64

	for _, quantity := range quantities {
		adjustedToPortions = append(adjustedToPortions, (quantity / 2) * float64(portions))
	}

	return adjustedToPortions
}
