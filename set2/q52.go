package set2

import "math"

func Q52(n int) int {
	if n == 0 {
		return 0
	}

	for i := n; i < math.MaxInt; i++ {
		if i%5 == 0 {
			return i
		}
	}

	return 0
}
