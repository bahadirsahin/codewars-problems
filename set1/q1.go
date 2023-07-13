package set1

func Q1(str string) (count int) {
	count = 0

	for _, r := range str {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			count++
		}
	}

	return count
}
