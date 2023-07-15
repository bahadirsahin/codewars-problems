package set1

import "math"

func Q8(a, b []int) []int {
	for _, v := range b {
		for i := 0; i < len(a); i++ {
			if a[i] == v {
				a[i] = math.MinInt32
			}
		}
	}

	res := []int{}

	for _, v := range a {
		if v > math.MinInt32 {
			res = append(res, v)
		}
	}

	return res
}
