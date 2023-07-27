package set4

func Q119(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}

	if n <= 3 {
		return signature[:n]
	}

	a, b, c, z := signature[0], signature[1], signature[2], float64(0)
	res := []float64{a, b, c}

	for i := 0; i < n-len(signature); i++ {
		z = a + b + c
		a = b
		b = c
		c = z
		res = append(res, z)
	}

	return res
}
