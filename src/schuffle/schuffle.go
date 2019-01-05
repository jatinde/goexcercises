package schuffle

import (
	"math/rand"
)

type Schuffleable interface {
	Len() int
	Swap(i, j int)
}

func Schuffle(s Schuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - 1)
		s.Swap(i, j)
	}
}
