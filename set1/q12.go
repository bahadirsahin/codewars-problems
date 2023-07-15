package set1

func Q12(dna string) string {
	if len(dna) == 0 {
		return ""
	}

	m := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
	}

	res := []rune{}

	for _, c := range dna {
		res = append(res, m[c])
	}

	return string(res)
}
