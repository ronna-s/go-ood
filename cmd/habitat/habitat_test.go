package main

import (
	"testing"

	"github.com/ronnas/go-ood/pkg/animal"
	"github.com/stretchr/testify/assert"
)

func TestHabitat_Add(t *testing.T) {
	h := Habitat{}
	l := &animal.Lion{}
	assert.Equal(t, animal.Deer{}.MakeSound(), h.Add(animal.Deer{}))
	assert.Equal(t, animal.Deer{}.MakeSound(), h.Add(animal.Deer{}))
	assert.Equal(t, animal.Lion{}.MakeSound(), h.Add(l))
	assert.True(t, l.Full())
	assert.Equal(t, 2, len(h.Animals))
	assert.Equal(t, animal.Deer{}.MakeSound(), h.Add(animal.Deer{}))
	assert.Equal(t, 3, len(h.Animals))
}
