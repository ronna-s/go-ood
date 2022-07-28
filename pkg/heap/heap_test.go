package heap

import (
	"testing"

	"golang.org/x/exp/constraints"
)

type Ordered[T constraints.Ordered] struct {
	T T
}

type intHeap = Heap[Ordered[int]]

func (o1 Ordered[T]) Less(o2 Ordered[T]) bool {
	return o1.T < o2.T
}

func verify(t *testing.T, h intHeap, i int) {
	t.Helper()
	n := len(h)
	j1 := 2*i + 1
	j2 := 2*i + 2
	if j1 < n {
		if h[j1].Less(h[i]) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i].T, j1, h[j1].T)
			return
		}
		verify(t, h, j1)
	}
	if j2 < n {
		if h[j2].Less(h[i]) {
			t.Errorf("heap invariant invalidated [%d] = %d > [%d] = %d", i, h[i].T, j1, h[j2].T)
			return
		}
		verify(t, h, j2)
	}
}

func TestHeap(t *testing.T) {
	var h intHeap

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
		x := h.Pop().T
		if i < 20 {
			h.Push(Ordered[int]{20 + i})
		}
		verify(t, h, 0)
		if x != i {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}
