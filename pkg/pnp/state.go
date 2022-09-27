package pnp

import (
	"math/rand"
)

// Rand - random function
var Rand = rand.Intn

//go:generate stringer -type=State

const (
	Calm State = iota
	SlightlyAnnoyed
	VeryAnnoyed
	Enraged
	Legacy
)

// React returns XP gained, Health gained and the new Production state
func (s State) React(a Action) (int, int, State) {
	chances := map[Action]int{
		DuckTyping:      70,
		TypeSafety:      80,
		Inheritance:     70,
		Interfaces:      80,
		Modules:         60,
		Generators:      80,
		MetaProgramming: 40,
		Generics:        80,
		DarkMagic:       20,
		Boredom:         80,
	}
	chanceOfSuccess, ok := chances[a]
	if !ok {
		panic("unexpected action")
	}
	var success bool
	success = Rand(100) > chanceOfSuccess
	switch s {
	case Calm:
		xp := Rand(10*int(a)/2+1) + 1
		if success {
			return xp, -Rand(10) - 1, SlightlyAnnoyed
		}
		return xp, Rand(10) + 1, s
	case SlightlyAnnoyed:
		xp := Rand(10*int(a)/2+1) + 11
		if success {
			return xp, -Rand(20) - 1, VeryAnnoyed
		}
		return xp, Rand(10) + 1, Calm
	case VeryAnnoyed:
		xp := Rand(10*int(a)/2+1) + 11
		if success {
			return xp, -Rand(20) - 1, Enraged
		}
		return xp, Rand(10) + 1, SlightlyAnnoyed
	case Enraged:
		xp := Rand(10*int(a)/2+1) + 21
		if success {
			return xp, -Rand(50) - 1, Legacy
		}
		return xp, Rand(10) + 1, VeryAnnoyed
	case Legacy:
		chances[DarkMagic] = 90 //DarkMagic is surprisingly effective against legacy
		chanceOfSuccess, ok = chances[a]
		if !ok {
			panic("unexpected action")
		}
		xp := Rand(10*int(a)/2+1) + 31
		if Rand(100) > chanceOfSuccess {
			return xp, -100, Legacy
		}
		return xp, Rand(10) + 1, Enraged
	}
	panic("unexpected state")
}
