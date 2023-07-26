package set4

import "sort"

func Q117(data []int) []int {
	if len(data) == 0 {
		return []int{}
	}

	res := []int{}
	res = append(res, data...)
	sort.Ints(res)
	rem := []int{}

	for i := len(res) - 2; i >= 0; i-- {
		rem = append(rem, res[i])
	}

	res = append(res, rem...)

	return res
}
