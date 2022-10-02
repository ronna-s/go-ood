package main

import (
	_ "embed"

	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnp/engine/tview"
)

func main() {
	//game := pnp.New(pnpdev.NewMinion(),pnpdev.NewGopher(), pnpdev.NewRubyist())
	game := pnp.New()
	game.Run(engine.New())
}
