package set5

import "strings"

func Q144(arr []string) string {
	if len(arr) == 0 {
		return ""
	}

	res, str := "", ""

	for i := 0; i < len(arr)-1; i++ {
		str = "("

		for j := 0; j < len(arr); j++ {
			if j == i {
				str += arr[j] + ", "
			} else {
				str += arr[j] + " "
			}
		}

		str += ")"
		str = strings.ReplaceAll(str, " )", ")")
		res += str
	}

	res = strings.ReplaceAll(res, ",)", ")")

	return res
}
