package set3

func Q88(a []int) int {
	if len(a) == 0 {
		return -1
	}

	m := map[int]int{}

	for _, v := range a {
		m[v]++
	}

	for k, v := range m {
		if v > (len(a) / 2) {
			return k
		}
	}

	return -1
}
