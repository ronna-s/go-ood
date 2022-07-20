package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRubyist(t *testing.T) {
	t.Run("abilities", func(t *testing.T) {
		r := Rubyist{Character: Character{X: 1}}
		cases := []Skill{DuckTyping, Module, DarkMagic}
		for i := 0; i < 5; i, r.X = i+1, r.X*10 {
			assert.ElementsMatch(t, cases[:i+1], r.Skills())
		}
	})
}
