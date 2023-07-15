package set1

func Q5(seq []int) int {
	if len(seq) == 0 {
		return 0
	}

	m := map[int]int{}

	for _, n := range seq {
		m[n]++
	}

	for k, v := range m {
		if v%2 == 1 {
			return k
		}
	}

	return 0
}
