package pnp

//go:generate stringer -type=Skill

// Skill represents a Player's skill in a Prorgammers & Platforms game
type Skill int

// Skills
const (
	DuckTyping Skill = iota
	TypeSafety
	Module
	Interface
	DarkMagic
	Generics
)
