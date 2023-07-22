package main

import (
	"codewars-problems/set3"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(
		set3.Q71(0),   // 1
		set3.Q71(1),   // 2
		set3.Q71(4),   // 70
		set3.Q71(7),   // 3432
		set3.Q71(13),  // 10400600
		set3.Q71(50),  // 100891344545564193334812497256
		set3.Q71(100), // 90548514656103281165404177077484163874504589675413336841320
	)
}
