package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

type (
// Gopher represents the P&P role of the Gopher
)

type Gopher struct {
	Character
}

func (g Rubyist) String() string {
	return "Rubyist"
}
func (g Gopher) String() string {
	return "Gopher"
}

var _ pnp.Player = &Gopher{}

// NewGopher ...
func NewGopher() *Gopher {
	return &Gopher{Character{H: 100}}
}

// Skills returns the list of abilities the Gopher has
// If XP is larger than 100 [TypeSafety, Interfaces, Generics]
// If XP is larger than 10 [TypeSafety, Interfaces]
// Anything else [TypeSafety]
func (g Gopher) Skills() []pnp.Skill {
	if g.XP() > 100 {
		return []pnp.Skill{pnp.TypeSafety, pnp.Interfaces, pnp.Generics}
	}
	if g.XP() > 10 {
		return []pnp.Skill{pnp.TypeSafety, pnp.Interfaces}
	}
	return []pnp.Skill{pnp.TypeSafety}
}

// Art renders the player's ascii art with the player state
func (g Gopher) Art() string {
	return fmt.Sprintf(gopherArt, g.H, g.X)
}
