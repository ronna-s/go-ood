package pnp

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

// Abilities returns the list of abilities the Rubyist has
func (r Rubyist) Skills() []Skill {
	abs := []Skill{DuckTyping}
	switch {
	case r.X >= 100:
		abs = append(abs, DarkMagic)
		fallthrough
	case r.X >= 10:
		abs = append(abs, Module)
	}
	return abs
}
