package set2

func Q55(numbers []int) bool {
	if len(numbers) < 2 {
		return true
	}

	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			return false
		}
	}

	return true
}
