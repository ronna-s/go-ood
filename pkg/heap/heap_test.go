package heap

import (
	"testing"
)

type Int int

func (i Int) Less(j Int) bool {
	return i < j
}

type intHeap = Heap[Int]

func verify(t *testing.T, h intHeap, i int) {
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
	var h intHeap

	verify(t, h, 0)
	for i := 20; i > 10; i-- {
		h.Push(Int(i))
	}

	verify(t, h, 0)
	for i := 10; i > 0; i-- {
		h.Push(Int(i))
		verify(t, h, 0)
	}

	for i := 1; len(h) > 0; i++ {
		x := h.Pop()
		if i < 20 {
			h.Push(Int(20 + i))
		}
		verify(t, h, 0)
		if x != Int(i) {
			t.Errorf("%d.th pop got %d; want %d", i, x, i)
		}
	}
}
