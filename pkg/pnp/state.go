package pnp

import (
	"math/rand"
)

// Rand - random function
var Rand = rand.Intn

// Calm ...
type Calm struct{}

// String ...
func (s Calm) String() string {
	return withColor(green, "calm")
}

// React returns XP gained, Health gained and the new Production state
func (s Calm) React(a Action) (int, int, State) {
	chances := map[Action]int{
		TypeSafety: 90,
		DuckTyping: 70,
		Interface:  90,
		Module:     50,
		Generics:   90,
		DarkMagic:  10,
	}
	xp := Rand(10*int(a)/2+1) + 1
	chanceOfSuccess, ok := chances[a]
	if !ok {
		panic("unexpected action")
	}

	if Rand(100) > chanceOfSuccess {
		return xp, -Rand(10) - 1, Annoyed{}
	}
	return xp, Rand(10) + 1, s
}

// Annoyed ...
type Annoyed struct {
	Very bool //How annoyed
}

// String ...
func (s Annoyed) String() string {
	if s.Very {
		return withColor(yellow, "very annoyed")
	}
	return withColor(yellow, "slightly annoyed")
}

// React returns XP gained, Health gained and the new Production state
func (s Annoyed) React(a Action) (int, int, State) {
	chances := map[Action]int{
		TypeSafety: 80,
		DuckTyping: 60,
		Interface:  80,
		Module:     40,
		Generics:   80,
		DarkMagic:  10,
	}
	xp := Rand(10*int(a)/2+1) + 11
	chanceOfSuccess, ok := chances[a]
	if !ok {
		panic("unexpected action")
	}

	if Rand(100) > chanceOfSuccess {
		if s.Very {
			return xp, -Rand(20) - 1, Enraged{}
		}
		return xp, -Rand(20) - 1, Annoyed{Very: true}
	}
	if s.Very {
		return xp, Rand(10) + 1, Annoyed{Very: false}
	}
	return xp, Rand(10) + 1, Calm{}
}

// Enraged is a product state is very eager to take a player's health in retaliation for unsuccessful actions
type Enraged struct{}

// String ...
func (s Enraged) String() string {
	return withColor(red, "Enraged!")
}

// React returns XP gained, Health gained and the new Production state
func (s Enraged) React(a Action) (int, int, State) {
	chances := map[Action]int{
		TypeSafety: 70,
		DuckTyping: 50,
		Interface:  70,
		Module:     30,
		Generics:   70,
		DarkMagic:  10,
	}
	xp := Rand(10*int(a)/2+1) + 21
	chanceOfSuccess, ok := chances[a]
	if !ok {
		panic("unexpected action")
	}
	if Rand(100) > chanceOfSuccess {
		return xp, -Rand(50) - 1, Legacy{}
	}

	return xp, Rand(10) + 1, Annoyed{Very: true}
}

// Legacy is a product state is very eager to take a player's health in retaliation for unsuccessful actions
type Legacy struct{}

// String ...
func (s Legacy) String() string {
	return withColor(purple, "LEGACY! - sudden death round")
}

// React returns XP gained, Health gained and the new Production state
func (s Legacy) React(a Action) (int, int, State) {
	// not a lot of chances against legacy - be very careful
	chances := map[Action]int{
		TypeSafety: 50,
		DuckTyping: 50,
		Interface:  50,
		Module:     30,
		Generics:   50,
		DarkMagic:  70, //dark magic is surprisingly effective in legacy mode
	}
	xp := Rand(10*int(a)/2+1) + 31
	chanceOfSuccess, ok := chances[a]
	if !ok {
		panic("unexpected action")
	}
	if Rand(100) > chanceOfSuccess {
		return xp, -100, Legacy{}
	}

	return xp, Rand(10) + 1, Enraged{}
}
