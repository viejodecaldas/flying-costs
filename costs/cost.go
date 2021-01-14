package costs

type FlyingCost struct {
	CityA	int
	CityB	int
}

// Function to find the minimum cost for an input array
func GetMinimumCost(input []FlyingCost) int {
	// Loop over the input array and compare flying costs and get the lowest
	// and sum to the previous cost gathered before
	var finalResult int
	for _, flyCost := range input {
		temp := flyCost.CityA
		if flyCost.CityA >= flyCost.CityB {
			temp = flyCost.CityB
		}
		finalResult += temp
	}

	return finalResult
}