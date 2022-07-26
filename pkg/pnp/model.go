package pnp

import "github.com/ronnas/go-ood/pkg/store"

type GameModel struct {
	store.Model[Game]
}

func (g GameModel) Cols() map[string]string {
	// type definition
	return map[string]string{
		"id":   "INTEGER PRIMARY KEY AUTOINCREMENT",
		"name": "varchar(10)",
	}
}
