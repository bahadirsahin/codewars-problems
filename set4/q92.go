package set4

func Q92(s string) string {
	if len(s) == 0 {
		return ""
	}

	arr := []rune{}

	for _, r := range s {
		if r != '#' {
			arr = append(arr, r)

			continue
		}

		if len(arr) == 0 {
			continue
		}

		arr = arr[:len(arr)-1]

	}

	return string(arr)
}
