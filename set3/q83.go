package set3

import (
	"crypto/md5"
	"fmt"
)

func Q83(str string) string {
	data := []byte(str)
	hash := fmt.Sprintf("%x", md5.Sum(data))

	return hash
}
