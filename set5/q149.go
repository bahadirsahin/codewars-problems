package set5

func Q149(board [3][3]int) int {
	// row
	for i := 0; i < 3; i++ {
		arr := []int{}

		for j := 0; j < 3; j++ {
			arr = append(arr, board[i][j])
		}

		if arr[0] == arr[1] &&
			arr[0] == arr[2] &&
			arr[1] == arr[2] {

			if arr[0] > 0 {
				return arr[0]
			}
		}
	}

	// col
	for i := 0; i < 3; i++ {
		arr := []int{}

		for j := 0; j < 3; j++ {
			arr = append(arr, board[j][i])
		}

		if arr[0] == arr[1] &&
			arr[0] == arr[2] &&
			arr[1] == arr[2] {

			if arr[0] > 0 {
				return arr[0]
			}
		}
	}

	// left diag
	arr := []int{}

	for i := 0; i < 3; i++ {
		arr = append(arr, board[i][i])
	}

	if arr[0] == arr[1] &&
		arr[0] == arr[2] &&
		arr[1] == arr[2] {

		if arr[0] > 0 {
			return arr[0]
		}
	}

	// right diag
	arr = []int{}

	for i := 0; i < 3; i++ {
		arr = append(arr, board[i][2-i])
	}

	if arr[0] == arr[1] &&
		arr[0] == arr[2] &&
		arr[1] == arr[2] {

		if arr[0] > 0 {
			return arr[0]
		}
	}

	// draw
	allPos := true

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				allPos = false
			}
		}
	}

	if allPos {
		return 0
	}

	// not finished
	return -1
}
