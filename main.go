package main

import (
	"codewars-problems/set1"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	arr1 := []int{98, 8, 72, -69, 78, -1, -69, 36, 34, -53, 24, 12, 44, -47, 8}
	arr2 := []int{-24, 1, 47, 64, 64, 144, 1156, 1296, 1936, 2809, 4761, 4761, 5184, 6084, 9604}

	spew.Dump(
		set1.Q25(arr1, arr2),
	)
}
