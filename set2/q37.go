package set2

import "math"

func Q37(arr []int) [2]int {
	if len(arr) == 0 {
		return [2]int{}
	}

	if len(arr) == 1 {
		return [2]int{arr[0], arr[0]}
	}

	min, max := math.MaxInt, math.MinInt

	for _, n := range arr {
		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	return [2]int{min, max}
}
