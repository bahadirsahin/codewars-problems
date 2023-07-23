package set4

func Q107(x, y, k uint64) uint64 {
	if x%k == 0 {
		return (y / k) - (x / k) + 1
	}

	return (y / k) - (x / k)
}
