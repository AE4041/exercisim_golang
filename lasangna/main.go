package lasangna

// Expected oven time to finish preparing the lasgna
const ovenTime = 40 // in minutes

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	if actualMinutesInOven < 0 {
		return 0
	}
	return ovenTime - actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if numberOfLayers < 0 {
		return 0
	}
	return numberOfLayers * 2
}

// ElapsedTime calculates the time elapsed cooking the lasagna.
// This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	if numberOfLayers < 0 {
		return 0
	}
	preparationTime := PreparationTime(numberOfLayers)
	return preparationTime + actualMinutesInOven
}
