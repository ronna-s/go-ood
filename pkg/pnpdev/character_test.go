package pnpdev

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacter(t *testing.T) {
	t.Run("XP and ApplyXPDiff", func(t *testing.T) {
		c := Character{}
		assert.Equal(t, 0, c.XP())
		c.ApplyXPDiff(10)
		assert.Equal(t, 10, c.XP())
	})
	t.Run("Health and ApplyHealthDiff", func(t *testing.T) {
		c := Character{}
		assert.Equal(t, 0, c.Health())
		gained := c.ApplyHealthDiff(10)
		assert.Equal(t, 10, c.Health())
		assert.Equal(t, 10, gained)
		t.Run("Health cannot exceed 100", func(t *testing.T) {
			gained := c.ApplyHealthDiff(1000)
			assert.Equal(t, 100, c.Health())
			assert.Equal(t, 90, gained)
		})
		gained = c.ApplyHealthDiff(-10)
		assert.Equal(t, -10, gained)
		assert.Equal(t, 90, c.Health())
		t.Run("Health cannot be negative", func(t *testing.T) {
			gained := c.ApplyHealthDiff(-1000)
			assert.Equal(t, -90, gained)
			assert.Equal(t, 0, c.Health())
		})

	})
	t.Run("Alive", func(t *testing.T) {
		c := Character{}
		assert.False(t, c.Alive())
		c.ApplyHealthDiff(20)
		assert.True(t, c.Alive())
		c.ApplyHealthDiff(-20)
		assert.False(t, c.Alive())
	})
}
