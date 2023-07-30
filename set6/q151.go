package set6

import (
	"math"
)

func Q146(arr []int, n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		sum := 0

		for _, v := range arr {
			sum += v
		}

		return sum
	}

	if n >= len(arr) {
		max := math.MinInt

		for _, v := range arr {
			if v > max {
				max = v
			}
		}

		return max
	}

	qtime := 0

	for {
		nums := []int{}

		for _, v := range arr {
			if v > 0 {
				nums = append(nums, v)

				if len(nums) == n {
					break
				}
			}
		}

		min := math.MaxInt

		for _, v := range nums {
			if v < min {
				min = v
			}
		}

		cnt := 0

		for i, v := range arr {
			if v > 0 {
				arr[i] -= min
				cnt++

				if cnt == len(nums) {
					break
				}
			}
		}

		qtime += min

		allZero := true

		for _, v := range arr {
			if v > 0 {
				allZero = false
			}
		}

		if allZero {
			break
		}
	}

	return qtime
}
