package set2

func Q48(start int, end int) int {
	cnt := 0
	flag := false

	for i := start; i <= end; i++ {
		flag = false

		n := i

		if n < 0 {
			n *= -1
		}

		for n > 0 {
			r := n % 10

			if r == 5 {
				flag = true
			}

			n /= 10
		}

		if !flag {
			cnt++
		}
	}

	return cnt
}
