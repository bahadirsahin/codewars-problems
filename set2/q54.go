package set2

func Q54(d, b int) int {
	if d >= b {
		return 0
	}

	max := 0

	for i := d; i <= b; i++ {
		if i%d == 0 {
			max = i
		}
	}

	return max
}
