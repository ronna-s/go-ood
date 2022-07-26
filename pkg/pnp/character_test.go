package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	t.Run("XP", func(t *testing.T) {
		c := Character{}
		c.X = 10
		assert.Equal(t, 10, c.XP())
	})
	t.Run("Health", func(t *testing.T) {
		c := Character{}
		c.H = 10
		assert.Equal(t, 10, c.Health())
	})
	t.Run("Alive", func(t *testing.T) {
		c := Character{}
		assert.False(t, c.Alive())
		c.H = 100
		c.Alive()
		assert.True(t, c.Alive())
	})
	t.Run("GainXP", func(t *testing.T) {
		c := Character{}
		c.GainXP(10)
	})
	t.Run("GainHealth", func(t *testing.T) {
		c := Character{}
		c.GainHealth(10)
		assert.Equal(t, 10, c.H)
		t.Run("health cannot exceed 100", func(t *testing.T) {
			gained := c.GainHealth(1000)
			assert.Equal(t, 100, c.Health())
			assert.Equal(t, 90, gained)
		})
		t.Run("health cannot be negative", func(t *testing.T) {
			c := Character{H: 10}
			gained := c.GainHealth(-1000)
			assert.Equal(t, -10, gained)
			assert.Equal(t, 0, c.Health())
		})
	})
}
