package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q132(0),          // "0.0.0.0"
		set5.Q132(32),         // "0.0.0.32"
		set5.Q132(2154959208), // "128.114.17.104"
	)
}
