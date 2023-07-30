package set6

func Q152(n int) int {
	if n == 0 {
		return 0
	}

	a, b, z := 1, 1, 0
	nums := []int{a, b}

	for i := 0; i < n-1; i++ {
		z = a + b
		a = b
		b = z
		nums = append(nums, z)
	}

	sum := 0

	for _, v := range nums {
		sum += 4 * v
	}

	return sum
}
