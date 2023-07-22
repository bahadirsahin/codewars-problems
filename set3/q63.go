package set3

func Q63(str string) bool {
	if len(str) == 0 {
		return true
	}

	m := map[rune]int{}

	for _, r := range str {
		m[r]++
	}

	for _, v := range m {
		if v > 1 {
			return false
		}
	}

	return true
}
