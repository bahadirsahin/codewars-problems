package set4

import "math"

func Q112(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}

	max := math.MinInt

	for i := 0; i < len(a1); i++ {
		x := a1[i]

		for j := 0; j < len(a2); j++ {
			y := a2[j]

			diff := 0
			xlen := len(x)
			ylen := len(y)

			if xlen < ylen {
				diff = ylen - xlen
			} else {
				diff = xlen - ylen
			}

			if diff > max {
				max = diff
			}
		}
	}

	return max
}
