package habitat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.Equal(t, "test", New("test").Name)
	assert.Equal(t, 0, len(New("test").Animals))
}

func TestHabitat_Add(t *testing.T) {
	h := Habitat{}
	h.Add(Deer{})
}
