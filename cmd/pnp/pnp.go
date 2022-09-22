package main

import (
	_ "embed"

	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnpdevs"
)

func main() {
	pnp.Run2(pnpdevs.NewRubyist(), pnpdevs.NewGopher())
}
