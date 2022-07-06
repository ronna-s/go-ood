package maze

import (
	"math/rand"
)

// Maze represents a 2D maze.
// The zero value is a very uninteresting case of a maze with no cells.
// Use New to generate a valid maze of the desired dimensions.
type Maze struct {
	// DimX and DimY are the dimensions of the maze
	DimX, DimY int
	// Cells is a slice of DimX * DimY maze cells, each cell has an array of two booleans represnting the existance of the passages Right and Down
	// If DimX is 5 then the x,y coordinates of cell of index 7 are 2,1 (index % DimX, index/DimY)
	// If the passage to the right of the cell is open then the first boolean is true
	// If the passage down from the cell is open then the second boolean is true.
	Cells [][2]bool
}

//Direction is an enum representing 4 possible values Up, Down, Left and Right
type Direction = int

// Valid Directions
const (
	Right Direction = iota
	Down
	Up
	Left
)

// Coords represents the X,Y coordinates of a cell in a maze
type Coords [2]int

const (
	// X Represents the X value of X,Y coordinates in a 2D maze
	X = iota
	// Y Represents the Y value of X,Y coordinates in a 2D maze
	Y
)

// Cell represents a cell in a grid or maze ranging from 0 to X*Y-1
type Cell = int

// CoordsFromCell takes a cell and returns its x,y coordinates
func (m Maze) CoordsFromCell(c Cell) Coords {
	return Coords{c % m.DimX, c / m.DimX}
}

// CellFromCoords takes x,y coordinates and returns the correrlating cell
func (m Maze) CellFromCoords(coords Coords) Cell {
	return m.DimX*coords[Y] + coords[X]
}

// represents the list of possible coordinates to add to a cell's coordinates to retrieve its neighbouring cells
var neighbours = make([]Coords, 4)

func init() {
	neighbours[Right] = Coords{1, 0}
	neighbours[Down] = Coords{0, 1}
	neighbours[Left] = Coords{-1, 0}
	neighbours[Up] = Coords{0, -1}
}

// Wall represents a wall between 2 cells in a 2D maze
type Wall struct {
	C, T Cell      // C for current Cell, T for target cell
	D    Direction // D for Direction
}

// Walls returns the valid walls of Cell c
func (m Maze) Walls(c Cell) (walls []Wall) {
	coords := m.CoordsFromCell(c)
	for d := len(neighbours) - 1; d >= 0; d-- {
		if neighbours[d][X]+coords[X] > -1 && neighbours[d][X]+coords[X] < m.DimX &&
			neighbours[d][Y]+coords[Y] > -1 && neighbours[d][Y]+coords[Y] < m.DimY {
			t := m.CellFromCoords(Coords{neighbours[d][X] + coords[X], neighbours[d][Y] + coords[Y]})
			w := Wall{c, t, d}
			walls = append(walls, w)
		}
	}
	return walls
}

// RemoveWall doesn't validate the existence of the wall to remove - use at own peril
func (m *Maze) RemoveWall(c Cell, d Direction) {
	switch d {
	case Up:
		m.Cells[c-m.DimX][Down] = true
	case Down:
		m.Cells[c][Down] = true
	case Left:
		m.Cells[c-1][Right] = true
	case Right:
		m.Cells[c][Right] = true
	}
}

// New generates a random maze using Prim's algorithm (non-simplified)
// The maze doesn't contain cycles, to generate a cycle remove any wall at random
func New(x, y int) Maze {
	m := Maze{x, y, make([][2]bool, x*y)}
	visited := make([]bool, x*y)
	s := 0
	visited[s] = true
	list := m.Walls(s)
	for len(list) != 0 {
		r := rand.Intn(len(list))
		w := list[r]
		list = append(list[:r], list[r+1:]...)
		if visited[w.C] != visited[w.T] {
			m.RemoveWall(w.C, w.D)
			if visited[w.C] {
				visited[w.T] = true
				list = append(list, m.Walls(w.T)...)
			} else {
				visited[w.C] = true
				list = append(list, m.Walls(w.C)...)
			}
		}
	}
	return m
}

// PathRight returns true if the path between the current cell and the cell to its right is open
func (m Maze) PathRight(cell Cell) bool {
	return m.Cells[cell][0]
}

// PathDown returns true if the path between the current cell and the cell below is open
func (m Maze) PathDown(cell Cell) bool {
	return m.Cells[cell][1]
}
