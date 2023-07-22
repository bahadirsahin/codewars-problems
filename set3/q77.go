package set3

func Q77(n int) string {
	if n == 0 {
		return "0"
	}

	if n == 1 {
		return "01"
	}

	f0, f1, res := "0", "01", ""

	for i := 1; i < n; i++ {
		res = f1 + f0
		f0 = f1
		f1 = res
	}

	return res
}
