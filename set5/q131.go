package set5

func Q131(nums []uint64) []uint64 {
	if len(nums) == 0 {
		return []uint64{0}
	}

	sum := uint64(0)

	for _, n := range nums {
		sum += n
	}

	res := []uint64{sum}

	for _, n := range nums {
		res = append(res, uint64(sum-n))
		sum -= n
	}

	return res
}
