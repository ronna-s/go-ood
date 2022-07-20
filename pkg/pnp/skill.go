//go:generate stringer -type=Skill
package pnp

// Skill represents a Player's skill in a Prorgammers & Platforms game
type Skill int

const (
	DuckTyping Skill = iota
	TypeSafety
	Module
	Interface
	DarkMagic
	Generics
)
