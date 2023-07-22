package set3

type Tuple struct {
	Char  rune
	Count int
}

func Q62(text string) []Tuple {
	if len(text) == 0 {
		return []Tuple{}
	}

	arr := []rune{}
	m := map[rune]int{}

	for _, r := range text {
		if m[r] == 0 {
			arr = append(arr, r)
		}

		m[r]++
	}

	res := []Tuple{}

	for _, r := range arr {
		t := Tuple{
			Char:  r,
			Count: m[r],
		}

		res = append(res, t)
	}

	return res
}
