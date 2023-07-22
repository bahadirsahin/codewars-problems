package set3

func Q69(n, m int) bool {
	if n == 0 || m == 0 {
		return false
	}

	min := 0

	if n < m {
		min = n
	} else {
		min = m
	}

	for i := 2; i <= min; i++ {
		if n%i == 0 && m%i == 0 {
			return false
		}
	}

	return true
}
