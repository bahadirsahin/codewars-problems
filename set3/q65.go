package set3

func Q65(n []int) []int {
	if len(n) == 0 {
		return []int{}
	}

	res := []int{}

	for i, v := range n {
		nn := ((i + 1) + v) % 10

		res = append(res, nn)
	}

	return res
}
