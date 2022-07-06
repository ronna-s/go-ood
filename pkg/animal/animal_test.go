package animal

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("test base animal", func(t *testing.T) {
		assert.Equal(t, "sound", BaseAnimal{"sound"}.MakeSound())
	})
	t.Run("test deer", func(t *testing.T) {
		d := NewDeer()
		t.Run("make sound", func(t *testing.T) {
			assert.Equal(t, "I'm so cute!", d.MakeSound())
		})
		t.Run("is an animal", func(t *testing.T) {
			var _ Animal = d
		})
	})
	t.Run("test lion", func(t *testing.T) {
		l := NewLion()
		t.Run("make sound", func(t *testing.T) {
			assert.Equal(t, "Roar!", l.MakeSound())
		})
		t.Run("is a predator", func(t *testing.T) {
			var _ Predator = l
		})
		t.Run("eats deer", func(t *testing.T) {
			assert.True(t, l.Eats(NewDeer()))
			assert.Equal(t, l.MakeSound(), l.Eat(nil))
			assert.True(t, l.Full())
			oldNow := Now
			Now = func() time.Time {
				return time.Now().Add(time.Hour)
			}
			assert.False(t, l.Full())
			Now = oldNow
		})
	})

}
