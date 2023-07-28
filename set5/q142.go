package set5

func Q142(n int) []int {
	if n == 1 {
		return []int{1}
	}

	res := []int{}

	for i := (n * n) - n; i < (n*n)+n; i++ {
		if i%2 == 1 {
			res = append(res, i)
		}
	}

	return res
}
