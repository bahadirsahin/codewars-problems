package set1

func Q24(n int) int {
	if n == 1 {
		return 1
	}

	first := 2*((n*(n-1))/2) + 1
	n++
	last := 2*((n*(n-1))/2) + 1
	sum := 0

	for i := first; i < last; i++ {
		if i%2 == 0 {
			continue
		}

		sum += i
	}

	return sum
}
