package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q131([]uint64{1, 2, 3, 4, 5, 6}),
		// []uint64{21, 20, 18, 15, 11, 6, 0}
	)
}
