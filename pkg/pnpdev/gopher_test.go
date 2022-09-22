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
		g := Gopher{Character: Character{X: 1}}
		cases := []pnp.Skill{pnp.TypeSafety, pnp.Interface, pnp.Generics}
		for i := 0; i < len(cases); i, g.X = i+1, g.X*10+1 {
			assert.ElementsMatch(t, cases[:i+1], g.Skills())
		}
	})
	t.Run("Art renders the player's ascii art with the player state", func(t *testing.T) {
		oldGopherArt := gopherArt
		defer func() { gopherArt = oldGopherArt }()
		gopherArt = "Nice Art [Health=%d,XP=%d]"
		g := Gopher{Character: Character{X: 10, H: 20}}
		assert.Equal(t, "Nice Art [Health=20,XP=10]", g.Art())

	})
}
