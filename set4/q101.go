package set4

func Q101(n int) int {
	if n == 0 {
		return 0
	}

	i := 2

	for {
		h := (i * (i - 1)) / 2

		if h >= n {
			break
		}

		i++
	}

	return i
}
