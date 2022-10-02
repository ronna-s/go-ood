package pnp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReactSuccessful(t *testing.T) {
	tt := []struct {
		State      State
		Health, XP int
	}{
		{State: Calm, Health: 1, XP: 1},
		{State: Annoyed, Health: 1, XP: 11},
		{State: Enraged, Health: 1, XP: 21},
		{State: Legacy, Health: 1, XP: 31},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%s", tc.State), func(t *testing.T) {
			expectedState := func() State {
				if i == 0 { //handle the case calm -> calm
					return tc.State
				}
				return tt[i-1].State
			}()
			oldRand := Rand
			defer func() {
				Rand = oldRand
			}()
			Rand = func(n int) int { return 0 }
			actions := []Action{
				Banana,
				DuckTyping,
				TypeSafety,
				Inheritance,
				Interfaces,
				Modules,
				Reflect,
				MetaProgramming,
				Generics,
				DarkMagic,
				Boredom,
			}
			for _, action := range actions {
				t.Run(action.String(), func(t *testing.T) {
					xp, h, state := tc.State.React(action)
					assert.Equal(t, tc.XP, xp)
					assert.Equal(t, tc.Health, h)
					assert.Equal(t, expectedState, state)
				})
			}
		})
	}
}
func TestReactUnsuccessful(t *testing.T) {
	tt := []struct {
		State      State
		Health, XP int
	}{
		{State: Calm, Health: -101, XP: 101},
		{State: Annoyed, Health: -101, XP: 111},
		{State: Enraged, Health: -101, XP: 121},
		{State: Legacy, Health: -100, XP: 131},
	}
	for i, tc := range tt {
		t.Run(fmt.Sprintf("%s", tc.State), func(t *testing.T) {
			expectedState := func() State {
				if i == len(tt)-1 { //handle the case calm -> calm
					return tc.State
				}
				return tt[i+1].State
			}()
			oldRand := Rand
			defer func() {
				Rand = oldRand
			}()
			Rand = func(n int) int { return 100 }
			actions := []Action{DuckTyping, TypeSafety, Modules, Interfaces, DarkMagic, Generics}
			for _, action := range actions {
				t.Run(action.String(), func(t *testing.T) {
					xp, h, state := tc.State.React(action)
					assert.Equal(t, tc.XP, xp)
					assert.Equal(t, tc.Health, h)
					assert.Equal(t, expectedState, state)
				})
			}
		})
	}
}
