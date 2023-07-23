package set4

func Q99(lst []int) []int {
	if len(lst) == 0 {
		return []int{}
	}

	res := []int{}

	for i := len(lst) - 1; i >= 0; i-- {
		res = append(res, lst[i])
	}

	return res
}
