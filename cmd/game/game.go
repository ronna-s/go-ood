package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ronnas/go-ood/pkg/pnp"
)

func main() {
	fmt.Println("New game started. A band of developers will attempt to survive against Production!")
	fmt.Println("What is the name of your band?")
	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic("error reading band name")
	}
	//todo: check if not exists
	g := pnp.NewGame(string(l), pnp.NewProduction(), pnp.NewRubyist(), pnp.NewGopher())
	g.Run()

}
