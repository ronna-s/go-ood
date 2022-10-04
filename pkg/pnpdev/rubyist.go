package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

type (
	// Rubyist represents the P&P role of the Rubyist
	Rubyist struct {
		Character
	}
)

var _ pnp.Player = NewRubyist()

// NewRubyist ...
func NewRubyist() *Rubyist {
	return &Rubyist{Character{H: 100}}
}

// AsciiArt renders the player's ascii art with the player state
func (r Rubyist) AsciiArt() string {
	return fmt.Sprintf(rubyistArt, r.H, r.X)
}

// Skills returns the list of abilities the Rubyist has
// If XP is larger than 100 [DuckTyping, Module, DarkMagic]
// If XP is larger than 10 [DuckTyping, Module, ]
// Anything else [DuckTyping]
func (r Rubyist) Skills() []pnp.Skill {
	skills := []pnp.Skill{pnp.DuckTyping, pnp.Inheritance, pnp.Modules, pnp.MetaProgramming, pnp.DarkMagic}
	if r.XP() > 100 {
		return skills
	}
	if r.XP() > 10 {
		return skills[:len(skills)-1]
	}
	return skills[:len(skills)-2]
}

func (g Rubyist) String() string {
	return "Rubyist"
}
