package set5

import (
	"fmt"
	"strconv"
	"strings"
)

func Q133(start, end string) int {
	return int(Ip2Int(end) - Ip2Int(start))
}

func Ip2Int(ip string) int64 {
	ts := strings.Split(ip, ".")
	bits := ""

	for _, t := range ts {
		d, _ := strconv.Atoi(t)
		bs := fmt.Sprintf("%08b", d)
		bits += bs
	}

	ipInt, _ := strconv.ParseInt(bits, 2, 64)

	return ipInt
}
