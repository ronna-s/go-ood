package main

import (
	_ "embed"

	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnpdev"
)

func main() {
	pnp.Run("pnp", pnpdev.NewGopher(), pnpdev.NewRubyist())
}
