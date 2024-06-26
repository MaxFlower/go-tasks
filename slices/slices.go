package slices

import (
	"fmt"
	"math/rand"
)

type Slices struct {
	origin []int
}

func (s *Slices) PrintOrigin() {
	fmt.Println(s.origin)
}

func (s *Slices) Fill(n int) {
	s.origin = make([]int, n)
	for i := 0; i < n; i++ {
		s.origin[i] = rand.Intn(9) + 1
	}
}

func (s *Slices) LeftTo(i int) []int {
	return s.origin[:i]
}

func (s *Slices) FromToRight(i int) []int {
	return s.origin[i:]
}
