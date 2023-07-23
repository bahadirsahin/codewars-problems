package set4

func Q91(str string) bool {
	if len(str) == 0 {
		return false
	}

	m := map[rune]rune{')': '(', ']': '[', '}': '{'}
	arr := []rune{}

	for _, r := range str {
		if r == '(' || r == '[' || r == '{' {
			arr = append(arr, r)
		}

		if r == ')' || r == ']' || r == '}' {
			if len(arr) == 0 {
				return false
			}

			if arr[len(arr)-1] != m[r] {
				return false
			}

			arr = arr[:len(arr)-1]
		}
	}

	return len(arr) == 0
}
