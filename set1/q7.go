package set1

func Q7(n int) int {
	for n > 9 {
		sum := 0
		nn := n

		for nn > 0 {
			sum += nn % 10
			nn /= 10
		}

		n = sum
	}

	return n
}
