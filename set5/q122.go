package set4

func Q122(prod uint64) [3]uint64 {
	if prod < 0 {
		return [3]uint64{0, 0, 0}
	}

	fib := []uint64{}
	a, b, p, s := uint64(0), uint64(1), uint64(0), uint64(0)

	for {
		fib = append(fib, s)

		if p == prod {
			ll := len(fib)
			return [3]uint64{fib[ll-3], fib[ll-2], 1}
		}

		if p > prod {
			ll := len(fib)
			return [3]uint64{fib[ll-3], fib[ll-2], 0}
		}

		s = a + b
		p = a * b
		a = b
		b = s
	}

	return [3]uint64{0, 0, 0}
}
