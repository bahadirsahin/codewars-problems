package set5

import (
	"fmt"
	"strconv"
)

func Q132(n uint32) string {
	if n == 0 {
		return "0.0.0.0"
	}

	bits := fmt.Sprintf("%032b", n)

	t1, t2, t3, t4 := bits[0:8], bits[8:16], bits[16:24], bits[24:32]

	t1d, _ := strconv.ParseInt(t1, 2, 64)
	t2d, _ := strconv.ParseInt(t2, 2, 64)
	t3d, _ := strconv.ParseInt(t3, 2, 64)
	t4d, _ := strconv.ParseInt(t4, 2, 64)

	return fmt.Sprintf("%d.%d.%d.%d", t1d, t2d, t3d, t4d)
}
