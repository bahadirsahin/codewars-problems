package set4

func Q109(numbers []int, target int) [2]int {
	if len(numbers) < 2 {
		return [2]int{}
	}

	idxMap := map[int]int{}
	subMap := map[int]int{}

	for i, n := range numbers {
		if subMap[n] == (target - n) {
			if i > idxMap[target-n] {
				return [2]int{idxMap[target-n], i}
			}

			return [2]int{i, idxMap[target-n]}
		}

		subMap[target-n] = n
		idxMap[n] = i
	}

	return [2]int{}
}
