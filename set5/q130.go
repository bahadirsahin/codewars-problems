package set5

import "math"

func Q130(m, n int) [][]int {
	if m == 0 || n == 0 {
		return [][]int{}
	}

	res := [][]int{}

	for i := m; i <= n; i++ {
		sqSum := 0

		for j := 1; j <= i; j++ {
			if i%j == 0 {
				sqSum += (j * j)
			}
		}

		sqCand := int(math.Sqrt(float64(sqSum)))

		if (sqCand * sqCand) == sqSum {
			arr := []int{i, sqSum}
			res = append(res, arr)
		}
	}

	return res
}
