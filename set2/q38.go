package set2

import "math"

func Q38(m int) int {
	// 1^3 + ... + n^3 = ((n * (n + 1)) / 2) ^ 2
	sqr := int(math.Sqrt(float64(m)))

	if (sqr * sqr) != m {
		return -1
	}

	// (n * (n + 1)) / 2) = sqr
	x := sqr
	x *= 2

	// n^2 + n - x = 0
	for i := 0; i < (x * x); i++ {
		if (i * (i + 1)) == x {
			return i
		}
	}

	return -1
}
