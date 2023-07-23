package main

import (
	"codewars-problems/set4"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set4.Q92("abc#d##c"),
		set4.Q92("abc##d######"),
		set4.Q92("#######"),
		set4.Q92(""),
	)
}
