package set4

func Q116(L []int) bool {
	if len(L) == 0 {
		return false
	}

	sum := 0

	for _, v := range L {
		sum += v
	}

	return sum > 0
}
