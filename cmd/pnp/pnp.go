package main

import (
	_ "embed"

	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnp/engine/tview"
	"github.com/ronna-s/go-ood/pkg/pnpdev"
)

func main() {
	game := pnp.New(pnpdev.NewGopher(), pnpdev.NewRubyist())
	game.Run(engine.New())
}
