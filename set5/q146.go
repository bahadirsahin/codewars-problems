package set5

import "strconv"

func Q146(a, b uint64) []uint64 {
	if a > b {
		return []uint64{}
	}

	res := []uint64{}

	for i := a; i <= b; i++ {
		str := strconv.FormatUint(i, 10)
		sum := uint64(0)

		for j, r := range str {
			s := string(r)
			ii, err := strconv.Atoi(s)

			if err != nil {
				continue
			}

			mul := uint64(1)

			for k := 0; k < j+1; k++ {
				mul *= uint64(ii)
			}

			sum += mul
		}

		if i == sum {
			res = append(res, i)
		}
	}

	return res
}
