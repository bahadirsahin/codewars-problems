package set3

func Q89(s int, g int) []int {
	if s%g != 0 {
		return []int{-1, -1}
	}

	div := s / g

	return []int{1 * g, (div - 1) * g}
}
