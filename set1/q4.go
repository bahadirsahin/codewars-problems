package set1

func Q4(num int) int {
	if num < 0 {
		return 0
	}

	sum := 0

	for i := 0; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
