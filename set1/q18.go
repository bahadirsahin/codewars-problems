package set1

func Q18(n int) int {
	if n < 10 {
		return 0
	}

	cnt := 0

	for n >= 10 {
		nn := n

		mul := 1

		for nn > 0 {
			mul *= nn % 10
			nn /= 10
		}

		n = mul
		cnt++
	}

	return cnt
}
