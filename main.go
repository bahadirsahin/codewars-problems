package main

import (
	"codewars-problems/set2"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}

	spew.Dump(
		set2.Q36(a1, a2),
	)
}
