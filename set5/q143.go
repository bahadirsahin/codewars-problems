package set5

func Q143(x int) int {
	b := 1
	ans := x

	for x != 0 {
		cur := (x-1)*b + (b - 1)

		if DigitSum(cur) > DigitSum(ans) || (DigitSum(cur) == DigitSum(ans) && cur > ans) {
			ans = cur
		}

		x = x / 10
		b = b * 10
	}

	return ans
}

func DigitSum(a int) int {
	sm := 0

	for a != 0 {
		sm = sm + a%10
		a = a / 10
	}

	return sm
}
