package set4

func Q114(letters []rune) rune {
	if len(letters) == 0 {
		return 'z'
	}

	if len(letters) == 1 {
		return letters[0]
	}

	ll := 0

	for _, r := range letters {
		ll += (int(r) - 96)
	}

	if ll%26 == 0 {
		return 'z'
	}

	ll = 96 + (ll % 26)

	return rune(ll)
}
