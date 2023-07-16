package set2

func Q41(h, bounce, window float64) int {
	if h <= 0 {
		return -1
	}

	if bounce <= 0 || bounce >= 1 {
		return -1
	}

	if window >= h {
		return -1
	}

	cnt := 1
	h *= bounce

	for h > window {
		cnt += 2
		h *= bounce
	}

	return cnt
}
