package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeer(t *testing.T) {
	assert.Equal(t, "I'm so pretty", NewDeer().Sound)
}
