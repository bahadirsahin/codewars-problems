package set6

func Q159(bonus int, price float64) int {
	cans := int(float64(bonus) / price)
	n := 1

	for {
		sum := 1
		sum *= n
		sum *= (n + 1)
		sum *= (2*n + 1)
		sum /= 6

		if sum > cans {
			break
		}

		n++
	}

	return n - 1
}
