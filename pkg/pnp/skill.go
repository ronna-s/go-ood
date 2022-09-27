package pnp

//go:generate stringer -type=Skill

// Skill represents a Player's skill in a Programmers & Platforms game
type Skill int

// Skills
const (
	DuckTyping Skill = iota
	TypeSafety
	Inheritance
	Interfaces
	Modules
	Generators
	MetaProgramming
	Generics
	DarkMagic
	Boredom
)
