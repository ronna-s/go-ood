package pnpdev

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

func TestGopher(t *testing.T) {
	t.Run("NewGopher", func(t *testing.T) {
		g := NewGopher()
		assert.Equal(t, 0, g.XP())
		assert.Equal(t, 100, g.Health())
	})
	t.Run("Skills", func(t *testing.T) {
		g := NewGopher()
		skills := []pnp.Skill{pnp.TypeSafety, pnp.Interfaces, pnp.Reflect, pnp.Generics, pnp.Boredom}
		g.X = 1
		assert.ElementsMatch(t, skills[:3], g.Skills())
		g.X = 11
		assert.ElementsMatch(t, skills[:4], g.Skills())
		g.X = 101
		assert.ElementsMatch(t, skills[:5], g.Skills())
	})
	t.Run("AsciiArt renders the player's ascii art with the player state", func(t *testing.T) {
		oldGopherArt := gopherArt
		defer func() { gopherArt = oldGopherArt }()
		gopherArt = "Nice Art [Health=%d,XP=%d]"
		g := Gopher{Character: Character{X: 10, H: 20}}
		assert.Equal(t, "Nice Art [Health=20,XP=10]", g.AsciiArt())
	})
}
