package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRubyist(t *testing.T) {
	t.Run("NewRubyist", func(t *testing.T) {
		r := NewRubyist()
		assert.Equal(t, 0, r.XP())
		assert.Equal(t, 100, r.Health())
	})
	t.Run("Skills", func(t *testing.T) {
		r := Rubyist{Character: Character{X: 1}}
		cases := []Skill{DuckTyping, Module, DarkMagic}
		for i := 0; i < len(cases); i, r.X = i+1, r.X*10+1 {
			assert.ElementsMatch(t, cases[:i+1], r.Skills())
		}
	})
	t.Run("Art renders the player's ascii art with the player's state", func(t *testing.T) {
		oldRubyistArt := rubyistArt
		defer func() { rubyistArt = oldRubyistArt }()
		rubyistArt = "Nice Art [Health=%d,XP=%d]"
		r := Rubyist{Character: Character{X: 10, H: 20}}
		assert.Equal(t, "Nice Art [Health=20,XP=10]", r.Art())

	})

}
