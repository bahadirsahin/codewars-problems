package set4

func Q113(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	m := map[int]bool{}

	for _, v := range arr {
		m[v] = true
	}

	for k, _ := range m {
		if !m[k*-1] {
			return k
		}
	}

	return 0
}
