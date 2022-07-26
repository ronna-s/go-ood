package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Record struct {
	Field1 int
	Field2 string
}

func (r Record) Schema() []SchemaEntry {
	return []SchemaEntry{
		{"f1", "int", &r.Field1},
	}
}
func TestStore(t *testing.T) {
	s, err := NewStore("records", func() Record { return Record{} })
	assert.NoError(t, err)
}
