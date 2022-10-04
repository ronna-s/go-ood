package pnpdev

import (
	"fmt"
	"github.com/ronna-s/go-ood/pkg/pnp"
)

// Minion represents a minion P&P player
// The zero value is a dead minion player.
type Minion struct {
	X, H int //X=XP, H=Health
}

// NewMinion returns a minion with 100 Health and 0 XP
func NewMinion() *Minion {
	return &Minion{H: 100}
}

// Alive checks if the player is (still) alive
func (m Minion) Alive() bool {
	return m.H > 0
}

// Health returns the player's health level
func (m Minion) Health() int {
	return m.H
}

// XP returns the player's xp level
func (m Minion) XP() int {
	return m.X
}

// ApplyXPDiff adds the given xp to the player's xp down to a minimum of 0
func (m *Minion) ApplyXPDiff(xp int) int {
	sum := m.X + xp
	if sum < 0 {
		xp = -m.X
		m.X = 0
		return xp
	}
	m.X = sum
	return xp
}

// ApplyHealthDiff adds the given health to the player's health down to a minimum of 0 and up to 100
func (m *Minion) ApplyHealthDiff(health int) int {
	sum := m.H + health
	if sum > 100 {
		health = 100 - m.H
		m.H = 100
		return health
	}
	if sum < 0 {
		health = -m.H
		m.H = 0
		return health
	}
	m.H = sum
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
