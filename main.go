package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q143(100),         // 99
		set5.Q143(2090),        // 1999
		set5.Q143(10193497399), // 1999
	)
}
