package main

import (
	"codewars-problems/set1"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	arr := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}

	spew.Dump(
		set1.Q5(arr),
	)
}
