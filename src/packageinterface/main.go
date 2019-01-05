package main

import (
	"fmt"
	"schuffle"
)

type intSlices []int

func (is intSlices) Len() int {
	return len(is)
}

func (is intSlices) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	is := intSlices{1, 3, 4, 2, 5, 7}
	schuffle.Schuffle(is)
	fmt.Printf("%v\n", is)
}
