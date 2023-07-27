package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q134("1.2.3.4"),         // true
		set5.Q134("123.45.67.89"),    // true
		set5.Q134("1.2.3"),           // false
		set5.Q134("123.456.78.90"),   // false
		set5.Q134("123.045.067.089"), // false
	)
}
