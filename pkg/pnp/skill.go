package pnp

//go:generate stringer -type=Skill

// Skill represents a Player's skill in a Programmers & Platforms game
type Skill int

// Skills
const (
	Banana Skill = iota
	DuckTyping
	TypeSafety
	Inheritance
	Interfaces
	Modules
	Reflect
	MetaProgramming
	Generics
	DarkMagic
	Boredom
)
