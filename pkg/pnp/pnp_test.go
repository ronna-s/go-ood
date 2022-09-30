package pnp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrToSentence(t *testing.T) {
	assert.Equal(t, "Hi There!", spaceCamelcase("HiThere!"))
	assert.Equal(t, "A A A", spaceCamelcase("AAA"))
}

type player struct {
	Player
	alive bool
}

func (p player) Alive() bool {
	return p.alive
}
func TestNextLivingPlayer(t *testing.T) {
	assert.Equal(t, -1, NextLivingPlayer([]*PlayerArt{{Player: &player{nil, false}}}, 0))
	assert.Equal(t, 0, NextLivingPlayer([]*PlayerArt{{Player: &player{nil, true}}}, 0))
	assert.Equal(t, 0, NextLivingPlayer([]*PlayerArt{{Player: &player{nil, true}}, {Player: &player{nil, false}}}, 0))
	assert.Equal(t, 0, NextLivingPlayer([]*PlayerArt{{Player: &player{nil, true}}, {Player: &player{nil, false}}}, 1))
}
