package main

import (
	"codewars-problems/set4"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set4.Q109([]int{1234, 5678, 9012}, 14690), // To(Equal([2]int{1, 2}))
	)
}
