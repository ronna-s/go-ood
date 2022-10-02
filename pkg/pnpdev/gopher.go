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

func (g Gopher) String() string {
	return "Gopher"
}

// NewGopher ...
func NewGopher() *Gopher {
	return &Gopher{Character{H: 100}}
}

// Skills returns the list of abilities the Gopher has
// If XP is larger than 100 [TypeSafety, Interfaces, Generics]
// If XP is larger than 10 [TypeSafety, Interfaces]
// Anything else [TypeSafety]
func (g Gopher) Skills() []pnp.Skill {
	skills := []pnp.Skill{pnp.TypeSafety, pnp.Interfaces, pnp.Reflect, pnp.Generics, pnp.Boredom}
	if g.XP() > 100 {
		return skills
	}
	if g.XP() > 10 {
		return skills[:len(skills)-1]
	}
	return skills[:len(skills)-2]
}

// Art renders the player's ascii art with the player state
func (g Gopher) Art() string {
	return fmt.Sprintf(gopherArt, g.H, g.X)
}
