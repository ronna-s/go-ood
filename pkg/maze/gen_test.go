package maze

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaze_RemoveWall(t *testing.T) {
	const x, y = 5, 3
	m := Maze{x, y, make([][2]bool, x*y)}
	assert.False(t, m.PathDown(1))
	m.RemoveWall(6, Up)
	assert.True(t, m.PathDown(1))
	assert.False(t, m.PathRight(5))
	m.RemoveWall(6, Left)
	assert.True(t, m.PathRight(5))
	assert.False(t, m.PathRight(6))
	m.RemoveWall(6, Right)
	assert.True(t, m.PathRight(6))
	assert.False(t, m.PathDown(6))
	m.RemoveWall(6, Down)
	assert.True(t, m.PathDown(6))
}
func TestMaze_CellIDFromCoords(t *testing.T) {
	assert.Equal(t, 0, New(1, 1).CellFromCoords(Coords{0, 0}))
	assert.Equal(t, 14, New(5, 7).CellFromCoords(Coords{4, 2}))
	assert.Equal(t, 22, New(5, 7).CellFromCoords(Coords{2, 4}))

	check := func(m Maze, x int, y int, z int) {
		for i := 0; i < m.DimX; i++ {
			for j := 0; j < m.DimY; j++ {
				if i == x && j == y {
					assert.Equal(t, z, m.CellFromCoords(Coords{x, y}))
					coords := m.CoordsFromCell(z)
					assert.Equal(t, coords[X], x)
					assert.Equal(t, coords[Y], y)
					return
				}
			}
		}
		t.Fatal("bad test params")
	}
	check(New(200, 33), 23, 2, 423)
	check(New(200, 33), 0, 0, 0)
	check(New(200, 33), 0, 1, 200)
	check(New(1, 1), 0, 0, 0)
	check(New(1, 2), 0, 1, 1)
}

func TestMaze_Down(t *testing.T) {
	const x, y = 5, 3
	m := Maze{x, y, make([][2]bool, x*y)}
	assert.False(t, m.PathDown(0))
	m.Cells[0][Down] = true
	assert.True(t, m.PathDown(0))
}

func TestMaze_Right(t *testing.T) {
	const x, y = 5, 3
	m := Maze{x, y, make([][2]bool, x*y)}
	assert.False(t, m.PathRight(0))
	m.Cells[0][Right] = true
	assert.True(t, m.PathRight(0))
}

func TestMaze_CellFromCoords(t *testing.T) {
	m := New(5, 3)
	assert.Equal(t, 0, m.CellFromCoords(Coords{0, 0}))
	assert.Equal(t, 1, m.CellFromCoords(Coords{1, 0}))
	assert.Equal(t, 5, m.CellFromCoords(Coords{0, 1}))
	assert.Equal(t, 6, m.CellFromCoords(Coords{1, 1}))
	assert.Equal(t, m.DimX*m.DimY-1, m.CellFromCoords(Coords{m.DimX - 1, m.DimY - 1}))
}
