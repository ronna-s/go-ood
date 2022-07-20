package travel

import (
	"testing"

	"github.com/ronnas/go-ood/pkg/maze"
	"github.com/stretchr/testify/assert"
)

func TestTravel(t *testing.T) {
	const x, y = 3, 2
	m := maze.Maze{DimX: x, DimY: y, Cells: make([][2]bool, x*y)}

	t.Run("travel down", func(t *testing.T) {
		tr := Travel{Dir: maze.Right, Maze: m}
		t.Run("out of bounds", func(t *testing.T) {
			assert.False(t, tr.up().CanGoDown())   // [0.-1]
			assert.False(t, tr.left().CanGoDown()) // [-1,0]
		})
		assert.False(t, tr.CanGoDown()) //[0,0]
		m.RemoveWall(m.CellFromCoords(tr.Coords), maze.Down)
		assert.True(t, tr.CanGoDown()) //[0,0]
		tr.TurnRight()
		assert.NoError(t, tr.Move())
		assert.Error(t, tr.Move())
		//go back up
		tr.TurnRight()
		assert.Error(t, tr.Move())
		tr.TurnRight()
		assert.NoError(t, tr.Move())
	})
	t.Run("travel right", func(t *testing.T) {
		tr := Travel{Dir: maze.Right, Maze: m}
		t.Run("out of bounds", func(t *testing.T) {
			assert.False(t, tr.up().CanGoRight())   // [0.-1]
			assert.False(t, tr.left().CanGoRight()) // [-1,0]
		})
		assert.False(t, tr.CanGoRight()) //[0,0]
		assert.Error(t, tr.Move())
		m.RemoveWall(m.CellFromCoords(tr.Coords), maze.Right)
		assert.True(t, tr.CanGoRight()) //[0,0]
		assert.NoError(t, tr.Move())
		assert.Error(t, tr.Move())
		//go back left
		tr.TurnRight()
		assert.Error(t, tr.Move())
		tr.TurnRight()
		assert.NoError(t, tr.Move())
	})
}
