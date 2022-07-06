package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	h := Heap[int]{3, 2, 1}
	h.Init()
	t.Run("test push", func(t *testing.T) {
		h.Push(2)
		h.Push(3)
		h.Push(-1)
		h.Push(4)
		s := h
		assert.Equal(t, -1, h.Pop())
		assert.Equal(t, -1, s[len(s)-1])
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 1, s[len(s)-2])
		assert.Equal(t, 2, h.Pop())
		assert.Equal(t, 2, s[len(s)-3])
		h.Push(5)
		assert.Equal(t, 2, h.Pop())
		assert.Equal(t, 3, h.Pop())
		h.Push(1)
		assert.Equal(t, 1, h.Pop())
		assert.Equal(t, 3, h.Pop())
		assert.Equal(t, 4, h.Pop())
		assert.Equal(t, 5, h.Pop())
	})

}
