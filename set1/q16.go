package set1

func Q16(a, b int) int {
	if a == b {
		return a
	}

	if a > b {
		tmp := a
		a = b
		b = tmp
	}

	sum := 0

	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}
