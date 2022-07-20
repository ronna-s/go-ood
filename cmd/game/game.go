package main

import (
	"github.com/ronnas/go-ood/pkg/pnp"
)

func main() {
	pnp.Game{Prod: pnp.NewProduction(), Players: []pnp.Player{pnp.NewRubyist(), pnp.NewGopher()}}.Run()
}
