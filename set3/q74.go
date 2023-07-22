package set3

func Q74(n, divisor int) int {
	cnt := 0

	for n/divisor > 0 {
		cnt++
		n /= divisor
	}

	return cnt
}
