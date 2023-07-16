package main

import (
	"codewars-problems/set2"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	arr := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}

	spew.Dump(
		set2.Q39(arr),
	)
}
