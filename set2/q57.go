package set2

func Q57(n int) int {
	if n < 2 {
		return 1
	}

	f := 1

	for i := 1; i <= n; i++ {
		f *= i
	}

	return f
}
