package main

import (
	"codewars-problems/set6"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set6.Q154(
			[]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}),
		// "-6,-3-1,3-5,7-11,14,15,17-20"
	)
}
