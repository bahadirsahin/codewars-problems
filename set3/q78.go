package set3

func Q78(n int) string {
	success, fail := "STRONG!!!!", "Not Strong !!"

	if n == 0 {
		return fail
	}

	sum, nn := 0, n

	for nn > 0 {
		d, f := nn%10, 1

		for i := 1; i <= d; i++ {
			f *= i
		}

		sum += f
		nn /= 10
	}

	if sum == n {
		return success
	}

	return fail
}
