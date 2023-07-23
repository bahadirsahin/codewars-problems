package set4

func Q97(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 1
	}

	a, b, c := 0, 1, 0

	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
	}

	return a
}
