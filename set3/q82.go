package set3

func Q82(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	oneCount := 0
	m := map[int]int{}

	for _, v := range arr {
		if v == 1 {
			oneCount++
		}

		m[v]++
	}

	return m[m[1]]
}
