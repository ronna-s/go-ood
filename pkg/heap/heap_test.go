package heap

import (
	"testing"

	"golang.org/x/exp/constraints"
)

//
//// Add takes any type with underlying type int
//// We can now increment all of those enums
//func Add[T ~int](i T, j T) T {
//	return i + j
//}
type Ordered[T constraints.Ordered] struct {
	T T
}

func (o1 Ordered[T]) Less(o2 Ordered[T]) bool {
	return o1.T < o2.T
}

var intHeap Heap[int]

func TestHeap(t *testing.T) {
	t.Run("")

}
