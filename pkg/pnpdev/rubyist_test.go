package pnpdev

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ronna-s/go-ood/pkg/pnp"
)

func TestRubyist(t *testing.T) {
	t.Run("NewRubyist", func(t *testing.T) {
		r := NewRubyist()
		assert.Equal(t, 0, r.XP())
		assert.Equal(t, 100, r.Health())
	})
	t.Run("Skills", func(t *testing.T) {
		r := Rubyist{Character: Character{X: 1}}
		skills := []pnp.Skill{pnp.DuckTyping, pnp.Inheritance, pnp.Modules, pnp.MetaProgramming, pnp.DarkMagic}
		r.X = 1
		assert.ElementsMatch(t, skills[:3], r.Skills())
		r.X = 11
		assert.ElementsMatch(t, skills[:4], r.Skills())
		r.X = 101
		assert.ElementsMatch(t, skills[:5], r.Skills())
	})

	t.Run("Art renders the player's ascii art with the player's state", func(t *testing.T) {
		oldRubyistArt := rubyistArt
		defer func() { rubyistArt = oldRubyistArt }()
		rubyistArt = "Nice Art [Health=%d,XP=%d]"
		r := Rubyist{Character: Character{X: 10, H: 20}}
		assert.Equal(t, "Nice Art [Health=20,XP=10]", r.Art())
	})
}
