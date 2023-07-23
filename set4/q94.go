package set4

import (
	"strconv"
	"strings"
	"unicode"
)

func Q94(str string) int {
	if len(str) == 0 {
		return 0
	}

	arr := []rune(str)

	for i, r := range arr {
		if !unicode.IsNumber(r) {
			arr[i] = ' '
		}
	}

	str = string(arr)
	ints := strings.Fields(str)
	sum := 0

	for _, v := range ints {
		i, _ := strconv.Atoi(v)

		sum += i
	}

	return sum
}
