package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGopher(t *testing.T) {
	t.Run("Skills", func(t *testing.T) {
		g := Gopher{Character: Character{X: 1}}
		cases := []Skill{TypeSafety, Interface, Generics}
		for i := 0; i < len(cases); i, g.X = i+1, g.X*10 {
			assert.ElementsMatch(t, cases[:i+1], g.Skills())
		}
	})
	t.Run("Image renders the player's ascii art with the player state", func(t *testing.T) {
		oldGopher := gopher
		g := Gopher{Character: Character{X: 10, H: 20}}
		gopher = "Nice Image [Health=%d,XP=%d]"
		defer func() { gopher = oldGopher }()
		assert.Equal(t, "Nice Image [Health=20,XP=10]", g.Image())

	})
}
