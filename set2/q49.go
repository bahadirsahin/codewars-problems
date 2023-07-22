package set2

import "sort"

func Q49(ages []int) [2]int {
	sort.Ints(ages)

	res := [2]int{
		ages[len(ages)-2],
		ages[len(ages)-1],
	}

	return res
}
