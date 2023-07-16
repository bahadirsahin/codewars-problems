package set2

import "strings"

func Q39(arr []string) []string {
	if len(arr) == 0 {
		return []string{}
	}

	m := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"EAST":  "WEST",
		"WEST":  "EAST",
	}

	res := []string{}

	for _, d := range arr {
		if len(res) == 0 {
			res = append(res, d)

			continue
		}

		last := res[len(res)-1]

		if strings.EqualFold(m[d], last) {
			res = res[:len(res)-1]

			continue
		}

		res = append(res, d)
	}

	return res
}
