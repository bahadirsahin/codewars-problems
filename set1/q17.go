package set1

func Q17(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}

	x, y := 0, 0

	for _, r := range walk {
		if r == 'n' {
			y++
		}

		if r == 'e' {
			x++
		}

		if r == 's' {
			y--
		}

		if r == 'w' {
			x--
		}
	}

	return (x == 0 && y == 0)
}
