package set5

func F(n int) int {
	if n == 0 {
		return 1
	}

	return n - M(F(n-1))
}

func M(n int) int {
	if n == 0 {
		return 0
	}

	return n - F(M(n-1))
}
