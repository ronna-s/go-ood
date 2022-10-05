package heap

import (
	"golang.org/x/exp/constraints"
	"testing"
)

type Ordered[T constraints.Ordered] struct {
	T T
}

func (o Ordered[T]) Less(o2 Ordered[T]) bool {
	return o.T < o2.T
}

type myHeap = Heap[Ordered[int]]

func verify(t *testing.T, h myHeap, i int) {
	t.Helper()
	n := len(h)
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h[j1].Less(h[i]) {
			t.Errorf("error")
			return
		}
		verify(t, h, j1)
	}
	if j2 < n {
		if h[j2].Less(h[i]) {
			t.Errorf("error")
			return
		}
		verify(t, h, j2)
	}
}

func TestHeap(t *testing.T) {
	var h myHeap

	verify(t, h, 0)
	for i := 20; i > 10; i-- {
		h.Push(Ordered[int]{i})
	}

	verify(t, h, 0)
	for i := 10; i > 0; i-- {
		h.Push(Ordered[int]{i})
		verify(t, h, 0)
	}

	for i := 1; len(h) > 0; i++ {
		x := h.Pop()
		if i < 20 {
			h.Push(Ordered[int]{20 + i})
		}
		verify(t, h, 0)
		y := Ordered[int]{i}
		if x != y {
			t.Errorf("%v.th pop got %v; want %v", i, x.T, i)
		}
	}
}
