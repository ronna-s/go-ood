package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ronnas/go-ood/pkg/heap"
	"golang.org/x/exp/constraints"
)

func main() {
	rand.Seed(time.Now().Unix())
	s := make([]int, 1000000)
	for i := range s {
		s[i] = rand.Intn(10000000)
	}

	fmt.Println(ExtractMin(s, 100))
}

// ExtractMin extract the min m items in h
func ExtractMin[T constraints.Ordered](s []T, m int) []T {
	h := heap.Heap[T](s)
	h.Init()
	if m > len(h) {
		m = len(h)
	}
	h.Init()

	for i := 0; i < m; i++ {
		h.Pop()
	}
	s = s[len(s)-m:]
	for i := 0; i < len(s)/2; i++ {
		t := s[len(s)-i-1]
		s[len(s)-i-1] = s[i]
		s[i] = t
	}
	return s
}
