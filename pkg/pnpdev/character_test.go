package pnpdev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	t.Run("XP and GainXP", func(t *testing.T) {
		c := Character{}
		assert.Equal(t, 0, c.XP())
		c.GainXP(10)
		assert.Equal(t, 10, c.XP())
	})
	t.Run("Health and GainHealth", func(t *testing.T) {
		c := Character{}
		assert.Equal(t, 0, c.Health())
		gained := c.GainHealth(10)
		assert.Equal(t, 10, c.Health())
		assert.Equal(t, 10, gained)
		t.Run("Health cannot exceed 100", func(t *testing.T) {
			gained := c.GainHealth(1000)
			assert.Equal(t, 100, c.Health())
			assert.Equal(t, 90, gained)
		})
		gained = c.GainHealth(-10)
		assert.Equal(t, -10, gained)
		assert.Equal(t, 90, c.Health())
		t.Run("Health cannot be negative", func(t *testing.T) {
			gained := c.GainHealth(-1000)
			assert.Equal(t, -90, gained)
			assert.Equal(t, 0, c.Health())
		})

	})
	t.Run("Alive", func(t *testing.T) {
		c := Character{}
		assert.False(t, c.Alive())
		c.GainHealth(20)
		assert.True(t, c.Alive())
		c.GainHealth(-20)
		assert.False(t, c.Alive())
	})
}
