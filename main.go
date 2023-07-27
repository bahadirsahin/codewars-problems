package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q133("10.0.0.0", "10.0.0.50"), // 50
		// set5.Q133("20.0.0.10", "20.0.1.0"), // 246
	)
}
