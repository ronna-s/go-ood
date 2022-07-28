// Package heap provides a generic heap
// By providing a generic heap we can avoid messy interface conversions and provide a friendly behavior
package heap

// Heap is a slice with heap properties of any type T with Less(T) bool function

// New takes a slice of type T with the method Less(T) and returns a new heap

// Init establishes the heap invariants required by the other routines in this package

// Push pushes the element x onto the heap.

// Pop removes and returns the minimum element (according to Less) from the heap.

// Remove removes and returns the element at index i from the heap.

// Fix re-establishes the heap ordering after the element at index i has changed its value.
