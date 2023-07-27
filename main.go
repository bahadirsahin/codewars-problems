package main

import (
	"codewars-problems/set5"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set5.Q135("1112223339"), // true
		set5.Q135("048665088X"), // true
		set5.Q135("1234554321"), // true
		set5.Q135("1234512345"), // false
		set5.Q135("ABCDEFGHIJ"), // false
		set5.Q135("XXXXXXXXXX"), // false
	)
}
