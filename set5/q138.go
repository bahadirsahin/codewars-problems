package set5

import (
	"math"
	"sort"
)

func Q138(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	freqMap := map[int]int{}

	for _, v := range nums {
		freqMap[v]++
	}

	maxOcc := math.MinInt

	for _, v := range freqMap {
		if v > maxOcc {
			maxOcc = v
		}
	}

	res := []int{}

	for k, v := range freqMap {
		if v == maxOcc {
			res = append(res, k)
		}
	}

	sort.Ints(res)

	return res[len(res)-1]
}
