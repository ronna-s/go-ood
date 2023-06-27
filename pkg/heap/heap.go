// Package heap provides a generic heap slice
// By providing a generic heap we can avoid messy interface conversions and provide a friendly interface
package heap

func (h Heap[T]) Init() {
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap[T]) Pop() T {
	n := len(*h) - 1
	h.Swap(0, n)
	h.down(0, n)
	old := *h
	n = len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Heap[T]) Push(t T) {
	*h = append(*h, t)
	h.up(len(*h) - 1)
}

func (h Heap[T]) Less(i, j int) bool {
	return h[i].Less(h[j])
}

func (h Heap[T]) Swap(i, j int) {
	tmp := h[i]
	h[i] = h[j]
	h[j] = tmp
}
func (h Heap[T]) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
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
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
