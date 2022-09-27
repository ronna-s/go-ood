// Package heap provides a generic heap
// By providing a generic heap we can avoid messy interface conversions and provide a friendly behavior
package heap

import (
	"sort"
)

type Heap[T interface{ Less(T) bool }] []T

func New[T interface{ Less(T) bool }]() *Heap[T] {
	return &Heap[T]{}
}

func (h Heap[T]) Init() {
	// heapify
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

func down[T interface{ Less(T) bool }](h Heap[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h[j2].Less(h[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h[i].Less(h[j]) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
func (h Heap[T]) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// The Interface type describes the requirements
// for a type using the routines in this package.
// Any type that implements it may be used as a
// min-heap with the following invariants (established after
// Init has been called or if the data is empty or sorted):
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// Note that Push and Pop in this interface are for package heap's
// implementation to call. To add and remove things from the heap,
// use heap.Push and heap.Pop.
type Interface interface {
	sort.Interface
	Push(x any) // add x as element Len()
	Pop() any   // remove and return element Len() - 1.
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(x T) {
	*h = append(*h, x)
	up(h, len(*h)-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func (h *Heap[T]) Pop() T {
	n := len(*h) - 1
	h.Swap(0, n)
	down(*h, 0, n)
	return h.pop()
}
func (h *Heap[T]) pop() T {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//
//// Remove removes and returns the element at index i from the heap.
//// The complexity is O(log n) where n = h.Len().
//func Remove(h Interface, i int) any {
//	n := h.Len() - 1
//	if n != i {
//		h.Swap(i, n)
//		if !down(h, i, n) {
//			up(h, i)
//		}
//	}
//	return h.Pop()
//}

// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
//func Fix(h Interface, i int) {
//	if !down(h, i, h.Len()) {
//		up(h, i)
//	}
//}

func up[T interface{ Less(T) bool }](h *Heap[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !(*h)[j].Less((*h)[i]) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

//
//func down(h Interface, i0, n int) bool {
//	i := i0
//	for {
//		j1 := 2*i + 1
//		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
//			break
//		}
//		j := j1 // left child
//		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
//			j = j2 // = 2*i + 2  // right child
//		}
//		if !h.Less(j, i) {
//			break
//		}
//		h.Swap(i, j)
//		i = j
//	}
//	return i > i0
//}
