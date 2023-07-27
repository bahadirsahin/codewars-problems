package set5

import "unicode"

func Q128(str string) bool {
	if len(str) == 0 {
		return false
	}

	for _, r := range str {
		if unicode.IsDigit(r) {
			continue
		}

		if unicode.IsLetter(r) {
			continue
		}

		return false
	}

	return true
}
