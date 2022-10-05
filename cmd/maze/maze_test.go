package main

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/ronna-s/go-ood/cmd/maze/mocks"
)

func TestSolveMaze(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var finished bool
	go func() {
		t.Run("test finished", func(t *testing.T) {
			g := mocks.Gopher{}
			g.On("Finished").Return(true)
			SolveMaze(&g)
			g.AssertExpectations(t)
		})

		// Test that the solution can solve any type of maze with any amount of dimensions
		t.Run("basic navigation", func(t *testing.T) {
			const (
				dims  = 15
				lefts = 12
			)
			g := mocks.Gopher{}
			nLefts := 0
			nRights := 0
			moved := false
			canMove := false
			g.On("Finished").Return(func() bool {
				return moved
			})

			g.On("TurnRight").Run(func(args mock.Arguments) {
				nRights++
				if nRights > nLefts && nRights-nLefts == dims-lefts {
					canMove = true
				} else if nLefts > nRights && nLefts-nRights == lefts {
					canMove = true
				} else {
					canMove = false
				}
			}).Maybe()
			g.On("TurnLeft").Run(func(args mock.Arguments) {
				nLefts++
				if nRights > nLefts && nRights-nLefts == 1 {
					canMove = true
				} else if nLefts > nRights && nLefts-nRights == 3 {
					canMove = true
				} else {
					canMove = false
				}
			}).Maybe()
			g.On("Move").Return(func() error {
				if !canMove {
					return errors.New("can't move")
				}
				moved = true
				return nil
			})

			SolveMaze(&g)
			g.AssertExpectations(t)
		})
		finished = true
		cancel()
	}()
	<-ctx.Done()
	if !finished {
		t.Fatal("test timed out")
	}
}
