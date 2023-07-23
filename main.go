package main

import (
	"codewars-problems/set4"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set4.Q107(6, 11, 2),    // 3
		set4.Q107(11, 345, 17), // 20
		set4.Q107(0, 1, 7),     // 1
		set4.Q107(20, 20, 2),   // 1
		set4.Q107(20, 20, 8),   // 0
		set4.Q107(19, 20, 2),   // 1
		set4.Q107(0, 10, 1),    // 11
		set4.Q107(11, 14, 2),   // 2
		set4.Q107(64, 73, 27),  // 0
	)
}
