package set6

import (
	"fmt"
	"strings"
)

func Q154(arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	res := ""
	nums := []int{}

	for _, n := range arr {
		if len(nums) == 0 {
			nums = append(nums, n)

			continue
		}

		last := nums[len(nums)-1]

		if (n - last) == 1 {
			nums = append(nums, n)

			continue
		}

		if len(nums) < 3 {
			for _, v := range nums {
				res += fmt.Sprintf("%d,", v)
			}

			nums = []int{n}

			continue
		}

		res += fmt.Sprintf("%d-%d,", nums[0], nums[len(nums)-1])

		nums = []int{n}
	}

	if len(nums) > 0 {
		if len(nums) < 3 {
			for _, v := range nums {
				res += fmt.Sprintf("%d,", v)
			}
		} else {
			res += fmt.Sprintf("%d-%d,", nums[0], nums[len(nums)-1])
		}
	}

	res = strings.TrimSuffix(res, ",")

	return res
}
