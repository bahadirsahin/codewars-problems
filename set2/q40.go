package set2

func Q40(chars []rune) rune {
	for i := 0; i < len(chars)-1; i++ {
		p, c := chars[i], chars[i+1]

		if (c - p) == 2 {
			return (c - 1)
		}
	}

	return ' '
}
