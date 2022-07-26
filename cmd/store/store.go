package main

import (
	"context"
	"fmt"

	"github.com/ronnas/go-ood/pkg/store"
)

type Thing struct {
	Field int
}

func (t *Thing) Schema() []store.SchemaEntry {
	return []store.SchemaEntry{
		{"field", "INTEGER", &t.Field}}
}

func main() {
	t := Thing{}
	s, _ := store.NewStore[*Thing]("things", func() *Thing { return &Thing{} })
	t.Field = 8
	var err error
	rec, err := s.Insert(context.Background(), &t)
	if err != nil {
		panic(err)
	}
	t = *rec.T
	fmt.Println(t, rec.ID, err)
	m, _ := s.Get(context.Background(), rec.ID)
	fmt.Println(m.ID, m.T.Field)
	//entries, err := s.All(context.Background())
	//fmt.Println(entries[0].T.Field)
	m.T.Field = 77
	s.Update(context.Background(), *m)
	m, _ = s.Get(context.Background(), rec.ID)
	fmt.Println(m.ID, m.T.Field)

}
