package set3

import "math"

func Q72(a []int, i int) int {
	n := a[i]
	least := math.MaxInt

	for _, v := range a {
		if v > n && v < least {
			least = v
		}
	}

	if least == math.MaxInt {
		return -1
	}

	for i, v := range a {
		if v == least {
			return i
		}
	}

	return 0
}
