package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGopher(t *testing.T) {
	t.Run("skills", func(t *testing.T) {
		g := Gopher{Character: Character{X: 1}}
		cases := []Skill{TypeSafety, Interface, Generics}
		for i := 0; i < 5; i, g.X = i+1, g.X*10 {
			assert.ElementsMatch(t, cases[:i+1], g.Skills())
		}
	})
}
