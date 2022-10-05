package pnpdev

import (
	"fmt"
	"github.com/ronna-s/go-ood/pkg/pnp"
)

type Character struct {
	X, H int //X=XP, H=Health
}

// Minion represents a minion P&P player
// The zero value is a dead minion player.
type Minion struct {
	Character
}

// NewMinion returns a minion with 100 Health and 0 XP
func NewMinion() *Minion {
	return &Minion{Character{H: 100}}
}

// Alive checks if the player is (still) alive
func (c Character) Alive() bool {
	return c.H > 0
}

// Health returns the player's health level
func (c Character) Health() int {
	return c.H
}

// XP returns the player's xp level
func (c Character) XP() int {
	return c.X
}

// ApplyXPDiff adds the given xp to the player's xp down to a minimum of 0
func (c *Character) ApplyXPDiff(xp int) int {
	sum := c.X + xp
	if sum < 0 {
		xp = -c.X
		c.X = 0
		return xp
	}
	c.X = sum
	return xp
}

// ApplyHealthDiff adds the given health to the player's health down to a minimum of 0 and up to 100
func (c *Character) ApplyHealthDiff(health int) int {
	sum := c.H + health
	if sum > 100 {
		health = 100 - c.H
		c.H = 100
		return health
	}
	if sum < 0 {
		health = -c.H
		c.H = 0
		return health
	}
	c.H = sum
	return health
}

// Skills returns the minion's skills which are Banana.
func (m Minion) Skills() []pnp.Skill {
	return []pnp.Skill{pnp.Banana}
}

// AsciiArt returns the minion's ascii-art
func (m Minion) AsciiArt() string {
	return fmt.Sprintf(minionArt, m.H, m.X)
}

func (m Minion) String() string {
	return "Minion"
}

type Gopher struct {
	Character
}

func NewGopher() *Gopher {
	return &Gopher{Character{H: 100}}
}

// Skills returns the minion's skills which are Banana.
func (g Gopher) Skills() []pnp.Skill {
	skills := []pnp.Skill{pnp.TypeSafety, pnp.Interfaces, pnp.Reflect, pnp.Generics, pnp.Boredom}
	if g.XP() <= 10 {
		return skills[:3]
	}
	if g.XP() <= 100 {
		return skills[:4]
	}
	return skills
}

// AsciiArt returns the minion's ascii-art
func (g Gopher) AsciiArt() string {
	return fmt.Sprintf(gopherArt, g.H, g.X)
}

type Rubyist struct {
	Character
}

func NewRubyist() *Rubyist {
	return &Rubyist{Character{H: 100}}
}

// Skills returns the minion's skills which are Banana.
func (r Rubyist) Skills() []pnp.Skill {
	skills := []pnp.Skill{pnp.DuckTyping, pnp.Inheritance, pnp.Modules, pnp.MetaProgramming, pnp.DarkMagic}
	if r.XP() <= 10 {
		return skills[:3]
	}
	if r.XP() <= 100 {
		return skills[:4]
	}
	return skills
}

// AsciiArt returns the minion's ascii-art
func (r Rubyist) AsciiArt() string {
	return fmt.Sprintf(rubyistArt, r.H, r.X)
}
