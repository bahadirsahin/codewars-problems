package set5

import (
	"strconv"
	"strings"
)

func Q135(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	for i := 0; i < 9; i++ {
		s := string(isbn[i])

		if strings.EqualFold(s, "X") {
			return false
		}

		_, err := strconv.Atoi(s)

		if err != nil {
			return false
		}
	}

	sum := 0

	for i := 0; i < 9; i++ {
		d, _ := strconv.Atoi(string(isbn[i]))
		sum += d * (i + 1)
	}

	if strings.EqualFold(string(isbn[9]), "X") {
		sum += 100
	} else {
		d, _ := strconv.Atoi(string(isbn[9]))
		sum += d * 10
	}

	return (sum%11 == 0)
}
