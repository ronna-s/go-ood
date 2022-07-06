package robot

import (
	"testing"

	"github.com/ronnas/go-ood/pkg/maze"
	"github.com/ronnas/go-ood/pkg/maze/travel"
	"github.com/stretchr/testify/assert"
)

func TestRobot_Finished(t *testing.T) {
	r := New(travel.New(maze.New(1, 1)))
	assert.True(t, r.Finished())
}
func TestRobot_Move(t *testing.T) {
	r := New(travel.New(maze.New(1, 1)))
	assert.Error(t, r.Move())
}

func TestRobot_TurnLeft(t *testing.T) {
	r := New(travel.New(maze.New(1, 1)))
	assert.Equal(t, 1, len(r.steps))
	r.TurnLeft()
	assert.Equal(t, 2, len(r.steps))
}

func TestRobot_TurnRight(t *testing.T) {
	r := New(travel.New(maze.New(1, 1)))
	assert.Equal(t, 1, len(r.steps))
	r.TurnRight()
	assert.Equal(t, 2, len(r.steps))
}
