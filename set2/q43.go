package set2

func Q43(n int) int {
	if n == 1 {
		return 1
	}

	cnt := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}

	return cnt
}
