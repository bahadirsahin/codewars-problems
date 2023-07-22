package set2

func Q60(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
