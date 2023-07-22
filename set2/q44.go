package set2

import "sort"

func Q44(numbers []int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	sort.Ints(numbers)

	return numbers
}
