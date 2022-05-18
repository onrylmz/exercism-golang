package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreprationTime int) int {
	if avgPreprationTime == 0 {
		avgPreprationTime = 2
	}

	return len(layers) * avgPreprationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		var layer = layers[i]

		if layer == "sauce" {
			sauce++
		} else if layer == "noodles" {
			noodles++
		}
	}

	noodles *= 50
	sauce *= .2

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) []string {
	myList = append(myList, friendsList[len(friendsList)-1])
	return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantites []float64 = make([]float64, len(quantities))
	for i, quantity := range quantities {
		quantity *= (float64(portions) / 2)
		scaledQuantites[i] = quantity
	}

	return scaledQuantites
}
