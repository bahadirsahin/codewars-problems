package set4

import (
	"math"
)

func Q111(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	allPos, allNeg := true, true

	for _, n := range nums {
		if n < 0 {
			allPos = false
		} else {
			allNeg = false
		}
	}

	if allNeg {
		return 0
	}

	if allPos {
		sum := 0

		for _, n := range nums {
			sum += n
		}

		return sum
	}

	max := math.MinInt

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			arr := []int{}

			for k := i; k <= j; k++ {
				arr = append(arr, nums[k])
			}

			sum := 0

			for _, n := range arr {
				sum += n
			}

			if sum > max {
				max = sum
			}
		}
	}

	return max
}
