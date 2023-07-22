package set2

import "math"

func Q51(array [3]int) int {
	min, max := math.MaxInt, math.MinInt

	for _, v := range array {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	for i, v := range array {
		if v != min && v != max {
			return i
		}
	}

	return 0
}
