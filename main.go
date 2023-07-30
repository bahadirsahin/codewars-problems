package main

import (
	"codewars-problems/set6"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set6.Q153([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
		// []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	)
}
