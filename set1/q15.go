package set1

func Q15(integers []int) int {
	evenArr, oddArr := []int{}, []int{}

	for _, v := range integers {
		if v%2 == 0 {
			evenArr = append(evenArr, v)
		} else {
			oddArr = append(oddArr, v)
		}

		if len(evenArr) >= 2 && len(oddArr) == 1 {
			return oddArr[0]
		}

		if len(evenArr) == 1 && len(oddArr) >= 2 {
			return evenArr[0]
		}
	}

	return 0
}
