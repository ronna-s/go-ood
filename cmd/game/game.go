package main

import (
	"github.com/ronnas/go-ood/pkg/pnp"
)

// ActionChooser ...
type ActionChooser func(player pnp.Player) pnp.Skill

// ChooseSkill ...
func (c ActionChooser) ChooseSkill(player pnp.Player) pnp.Skill {
	return c(player)
}

func main() {
	pnp.Game{Prod: pnp.NewProduction(), Players: []pnp.Player{pnp.NewRubyist(), pnp.NewGopher()}}.Run()
}
