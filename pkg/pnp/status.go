package pnp

import "math/rand"

// Calm ...
type Calm struct{}

// String ...
func (s Calm) String() string {
	return "calm"
}

// React returns XP gained, Health gained and the new Production status
func (s Calm) React(a Action) (int, int, Status) {
	xp := rand.Intn(10*int(a)/2+1) + 1
	chanceOfSuccess := func(a Action) int {
		switch a {
		case TypeSafety:
			return 90
		case DuckTyping:
			return 70
		case Module:
			return 50
		case Interface:
			return 90
		case DarkMagic:
			return 10
		case Generics:
			return 90
		}
		panic("unexpected action")
		return 0
	}(a)

	if rand.Intn(100) > chanceOfSuccess {
		return xp, -rand.Intn(10) - 1, Annoyed{}
	}
	return xp, rand.Intn(10) + 1, s
}

// Annoyed ...
type Annoyed struct {
	Very bool //How annoyed
}

// String ...
func (s Annoyed) String() string {
	if s.Very {
		return "very annoyed"
	}
	return "slightly annoyed"
}

// React returns XP gained, Health gained and the new Production status
func (s Annoyed) React(a Action) (int, int, Status) {
	xp := rand.Intn(10*int(a)/2+1) + 1
	chanceOfSuccess := func(a Action) int {
		switch a {
		case TypeSafety:
			return 90
		case DuckTyping:
			return 70
		case Module:
			return 50
		case Interface:
			return 90
		case DarkMagic:
			return 10
		case Generics:
			return 10
		}
		panic("unexpected action")
		return 0
	}(a)

	if rand.Intn(100) > chanceOfSuccess {
		if s.Very {
			return xp, -rand.Intn(20) - 1, Enraged{}
		}
		return xp, -rand.Intn(20) - 1, Annoyed{Very: true}
	}
	if s.Very {
		return xp, rand.Intn(10) + 1, Annoyed{Very: false}
	}
	return xp, rand.Intn(10) + 1, Calm{}
}

// Enraged is a product state is very eager to take a player's health in retaliation for unsuccessful actions
type Enraged struct{}

// String ...
func (s Enraged) String() string {
	return "enraged"
}

// React returns XP gained, Health gained and the new Production status
func (s Enraged) React(a Action) (int, int, Status) {
	xp := rand.Intn(10*int(a)/2+1) + 1
	chanceOfSuccess := func(a Action) int {
		switch a {
		case TypeSafety:
			return 90
		case DuckTyping:
			return 70
		case Module:
			return 50
		case Interface:
			return 90
		case DarkMagic:
			return 10
		case Generics:
			return 10
		}
		panic("unexpected action")
		return 0
	}(a)

	if rand.Intn(100) > chanceOfSuccess {
		return xp, -rand.Intn(50) - 1, Enraged{}
	}

	return xp, rand.Intn(10) + 1, Annoyed{Very: true}
}
