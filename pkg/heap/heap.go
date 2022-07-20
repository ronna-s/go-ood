package heap

import (
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] []T

// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func (h *Heap[T]) Init() {
	// heapify
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap[T]) Pop() T {
	n := len(*h) - 1
	t := (*h)[n]
	(*h)[n] = (*h)[0]
	(*h)[0] = t
	h.down(0, n)
	old := *h
	n = len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func (h Heap[T]) Remove(i int) T {
	n := len(h) - 1
	if n != i {
		t := h[n]
		h[n] = h[i]
		h[i] = t
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.Pop()
}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func (h Heap[T]) Fix(i int) {
	if !h.down(i, len(h)) {
		h.up(i)
	}
}

func (h Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(h[j] < h[i]) {
			break
		}
		t := h[j]
		h[j] = h[i]
		h[i] = t
		j = i
	}
}

func (h Heap[T]) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h[j2] < h[j1] {
			j = j2 // = 2*i + 2  // right child
		}
		if !(h[j] < h[i]) {
			break
		}
		t := h[j]
		h[j] = h[i]
		h[i] = t
		i = j
	}
	return i > i0
}
