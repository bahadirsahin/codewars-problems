package set2

func Q31(busStops [][2]int) int {
	res := 0

	for _, v := range busStops {
		res += v[0] - v[1]
	}

	return res
}
