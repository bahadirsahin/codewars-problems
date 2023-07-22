package set3

func Q90(x, y uint32) uint32 {
	if y == 0 {
		return x
	}

	return Q90(y, x%y)
}
