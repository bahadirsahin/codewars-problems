package set3

import "strings"

func Q67(grains <-chan string) (good int, bad int) {
	for w := range grains {
		if strings.EqualFold(w, "good") {
			good++
		}

		if strings.EqualFold(w, "bad") {
			bad++
		}
	}

	return
}
