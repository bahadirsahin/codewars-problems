package set4

import "unicode"

func Q106(word string) []string {
	if len(word) == 0 {
		return []string{}
	}

	res := []string{}

	for i := 0; i < len(word); i++ {
		if word[i] == ' ' {
			continue
		}

		arr := []rune{}

		for j, r := range word {
			if j != i {
				arr = append(arr, r)

				continue
			}

			arr = append(arr, unicode.ToUpper(r))
		}

		res = append(res, string(arr))
	}

	return res
}
