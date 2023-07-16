package set2

func Q33(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	for i := 0; i < len(arr); i++ {
		leftSum, rightSum := 0, 0

		for j := i - 1; j >= 0; j-- {
			leftSum += arr[j]
		}

		for j := i + 1; j <= len(arr)-1; j++ {
			rightSum += arr[j]
		}

		if leftSum == rightSum {
			return i
		}
	}

	return -1
}
