package set4

import "math"

func Q115(n int64) int64 {
	if n/10 == 0 {
		return n
	}

	arr := []int64{}
	nn := n

	for nn > 0 {
		arr = append([]int64{nn % int64(10)}, arr...)
		nn /= 10
	}

	nums := []int64{}

	for i := 0; i < len(arr); i++ {
		num := int64(0)

		for j := 0; j < len(arr); j++ {
			num += arr[j] * int64(math.Pow10(len(arr)-j-1))
		}

		nums = append(nums, num)

		tmp := arr[i]

		for j := i; j < len(arr)-1; j++ {
			arr[j] = arr[j+1]
		}

		arr[len(arr)-1] = tmp
	}

	max := int64(math.MinInt64)

	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	return max
}
