package pnpdev

import (
	"fmt"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

type Minion struct {
	X, H int
}

func NewMinion() *Minion {
	return &Minion{H: 100}
}

func (m Minion) Alive() bool {
	return m.H > 0
}

func (m *Minion) ApplyXPDiff(xp int) {
	m.X += xp
}

func (m *Minion) ApplyHealthDiff(health int) int {
	sumH := m.H + health
	if sumH > 100 {
		health = 100 - m.H
		m.H = 100
		return health
	}
	if sumH < 0 {
		health = 0 - m.H
		m.H = 0
		return health
	}
	m.H = sumH
	return health
}

func (m Minion) Skills() []pnp.Skill {
	return []pnp.Skill{pnp.Banana}
}

func (m Minion) Health() int {
	return m.H
}

func (m Minion) Art() string {
	return fmt.Sprintf(minionArt, m.H, m.X)

}

func (m Minion) XP() int {
	return m.X
}

func (m Minion) String() string {
	return "Minion"
}
