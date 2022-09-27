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

// Art renders the player's ascii art with the player state
func (r Rubyist) Art() string {
	return fmt.Sprintf(rubyistArt, r.H, r.X)
}

// Skills returns the list of abilities the Rubyist has
// If XP is larger than 100 [DuckTyping, Module, DarkMagic]
// If XP is larger than 10 [DuckTyping, Module, ]
// Anything else [DuckTyping]
func (r Rubyist) Skills() []pnp.Skill {
	s := make([]pnp.Skill, pnp.Boredom/2)
	for i := range s {
		s[i] = pnp.Skill(i * 2)
		i++
	}
	if r.XP() > 100 {
		return s
	}
	if r.XP() > 10 {
		return s[:4]
	}
	return s[:3]
}
