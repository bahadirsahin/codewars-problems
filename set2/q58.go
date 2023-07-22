package set2

func Q58(n int) int {
	if n == 1 {
		return 1
	}

	r := (n) * (n + 1)
	r /= 2
	r *= r

	return r
}
