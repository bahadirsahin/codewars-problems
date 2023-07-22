package set2

func Q47(n int, d int) int {
	if n == 0 {
		return 0
	}

	cnt := 0

	for k := 0; k <= n; k++ {
		kk := k * k

		for kk >= 0 {
			r := kk % 10

			if r == d {
				cnt++
			}

			kk /= 10

			if kk == 0 {
				break
			}
		}
	}

	return cnt
}
