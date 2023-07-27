package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q129(7775460), // "(2**2)(3**3)(5)(7)(11**2)(17)"
		set5.Q129(79340),   // "(2**2)(5)(3967)"
	)
}
