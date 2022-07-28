// Package travel provides travel for package maze
package travel

import (
	"errors"

	"github.com/ronnas/go-ood/pkg/maze"
)

// Travel represents a being that can travel through a maze
type Travel struct {
	Dir    maze.Direction
	Coords maze.Coords
	Maze   maze.Maze
}

// New returns a valid gopher that is ready to roam through a maze
func New(m maze.Maze) Travel {
	return Travel{Dir: maze.Right, Coords: maze.Coords{0, 0}, Maze: m}
}

// left returns the Coordinates of the cell to the left of the current cell
// The result coordinates may not be valid (out of bounds)
func (t Travel) left() Travel {
	t.Coords[maze.X] = t.Coords[maze.X] - 1
	return t
}

// right returns the Coordinates of the cell to the right of the current cell
// The result coordinates may not be valid (out of bounds)
func (t Travel) right() Travel {
	t.Coords[maze.X] = t.Coords[maze.X] + 1
	return t
}

// up returns the Coordinates of the cell above the current cell
// The result coordinates may not be valid (out of bounds)
func (t Travel) up() Travel {
	t.Coords[maze.Y] = t.Coords[maze.Y] - 1
	return t
}

// down returns the Coordinates of the cell below the current cell
// The result coordinates may not be valid (out of bounds)
func (t Travel) down() Travel {
	t.Coords[maze.Y] = t.Coords[maze.Y] + 1
	return t
}

// CanGoRight returns false if there's a wall between the current cell and the cell to the right or
// if the cell's coordinates are invalid
func (t Travel) CanGoRight() bool {
	if t.Coords[maze.Y] < 0 || t.Coords[maze.Y] >= t.Maze.DimY ||
		t.Coords[maze.X] < 0 || t.Coords[maze.X] >= t.Maze.DimX {
		return false
	}
	return t.Maze.PathRight(t.Maze.CellFromCoords(t.Coords))
}

// CanGoDown returns false if there's a wall between the current cell and the cell below it or
// if the cell's coordinates are invalid
func (t Travel) CanGoDown() bool {
	if t.Coords[maze.Y] < 0 || t.Coords[maze.Y] >= t.Maze.DimY ||
		t.Coords[maze.X] < 0 || t.Coords[maze.X] >= t.Maze.DimX {
		return false
	}
	return t.Maze.PathDown(t.Maze.CellFromCoords(t.Coords))
}

// Move tries to move the gopher to the cell in the direction that the gopher is facing and returns an error
// if there's no passage
func (t *Travel) Move() error {
	switch t.Dir {
	case maze.Right:
		if t.CanGoRight() {
			*t = t.right()
			return nil
		}
	case maze.Down:
		if t.CanGoDown() {
			*t = t.down()
			return nil
		}
	case maze.Left:
		if t.left().CanGoRight() {
			*t = t.left()
			return nil
		}
	case maze.Up:
		if t.up().CanGoDown() {
			*t = t.up()
			return nil
		}
	}
	return errors.New("invalid move")
}

func (t *Travel) TurnRight() {
	switch t.Dir {
	case maze.Right:
		t.Dir = maze.Down
	case maze.Down:
		t.Dir = maze.Left
	case maze.Left:
		t.Dir = maze.Up
	case maze.Up:
		t.Dir = maze.Right
	default:
		panic("something went wrong here")
	}
}
