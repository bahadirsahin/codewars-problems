package set1

import "sort"

func Q25(arr1 []int, arr2 []int) bool {
	if arr1 == nil || arr2 == nil {
		return false
	}

	if len(arr1) != len(arr2) {
		return false
	}

	arr := []int{}

	for _, n := range arr1 {
		arr = append(arr, n*n)
	}

	sort.Ints(arr)
	sort.Ints(arr2)

	for i, n := range arr {
		if n != arr2[i] {
			return false
		}
	}

	return true
}
