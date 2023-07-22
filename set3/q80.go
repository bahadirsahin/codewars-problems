package set3

import (
	"fmt"
	"strings"
)

func Q80(word string) string {
	if len(word) == 0 {
		return ""
	}

	if word[0] != word[len(word)-1] {
		return fmt.Sprintf("The %s", strings.Title(word))
	}

	res := strings.Title(word)
	res += word[1:]

	return res
}
