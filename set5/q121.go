package set5

import "math"

func Q121(recipe, available map[string]int) int {
	if len(recipe) == 0 || len(available) == 0 {
		return 0
	}

	for k, _ := range recipe {
		if available[k] == 0 {
			return 0
		}
	}

	cakes := math.MaxInt

	for k, v := range recipe {
		div := available[k] / v

		if div < cakes {
			cakes = div
		}
	}

	return cakes
}
