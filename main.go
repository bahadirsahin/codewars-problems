package main

import (
	"codewars-problems/set3"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set3.Q72([]int{1, 3, 5, 2, 4}, 0),
		set3.Q72([]int{4, 1, 3, 5, 6}, 4),
	)
}
