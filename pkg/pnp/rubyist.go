package pnp

import (
	_ "embed"
	"fmt"
	"sort"
)

type (
	// Rubyist represents the P&P role of the Rubyist
	Rubyist struct {
		Character
	}
)

// NewRubyist ...
func NewRubyist() *Rubyist {
	return &Rubyist{Character{H: 100, Name: "Rubyist"}}
}

//go:embed resources/rubyist.txt
var rubyist string

func (r Rubyist) Image() string {
	return fmt.Sprintf(rubyist, r.H, r.X)
}

// Skills returns the list of abilities the Rubyist has
func (r Rubyist) Skills() []Skill {
	abs := []Skill{DuckTyping}
	switch {
	case r.X >= 100:
		abs = append(abs, DarkMagic)
		fallthrough
	case r.X >= 10:
		abs = append(abs, Module)
	}
	sort.Slice(abs, func(i, j int) bool { return abs[i] < abs[j] })
	return abs
}
