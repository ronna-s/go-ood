package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ronnas/go-ood/pkg/pnp"
)

func main() {
	fmt.Println("What is the name of your band?")
	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic("error reading band name")
	}
	//todo: check if not exists
	g := pnp.NewGame(string(l), pnp.NewProduction(), pnp.NewRubyist(), pnp.NewGopher())
	fmt.Printf("New game started. The band of developers '%s' will attempt to survive against Production!\n", g.Name)
	g.Run()

}

//
//func selectGame(games []store.Model[pnp.Game]) store.Model[pnp.Game] {
//	for {
//		fmt.Println("Please choose the band you would like to play!:")
//		for i := range games {
//			fmt.Printf("[%d]%s\n", i+1, games[i].T.Name)
//		}
//		fmt.Print("> ")
//		var i int
//		if _, err := fmt.Scanln(&i); err != nil {
//			fmt.Printf("failed parsing input %s\n", err)
//		} else if i < 1 || i > len(games) {
//			fmt.Printf("invalid option %d\n", i)
//		} else {
//			return games[i-1]
//		}
//	}
//}
//
//func main() {
//	fmt.Print("[N]ew P&P game?\n")
//	fmt.Print("[L]oad game?\n")
//	for {
//		b, err := bufio.NewReader(os.Stdin).ReadByte()
//		if err != nil {
//			panic(err)
//		}
//		if b == 'N' {
//			fmt.Println("What is the name of your band?")
//			l, _, err := bufio.NewReader(os.Stdin).ReadLine()
//			if err != nil {
//				panic("error reading band name")
//			}
//			//todo: check if not exists
//			 = pnp.NewGame(string(l), pnp.NewProduction(), pnp.NewRubyist(), pnp.NewGopher())
//			fmt.Printf("New game started. The band of developers '%s' will attempt to survive against Production!\n", g.T.Name)
//			break
//		} else if b == 'L' {
//			games, _ := s.All(context.Background())
//			g = selectGame(games)
//			fmt.Printf("Resuming game. The band of developers '%s' will attempt to survive against Production!\n", g.T.Name)
//			break
//		} else {
//			fmt.Println("invalid option")
//		}
//	}
//
//	fmt.Println("Press enter to continue...")
//	_, _ = bufio.NewReader(os.Stdin).ReadByte()
//
//	g.T.Run()
//	fmt.Println("Saving game...")
//	if g.ID == 0 {
//		_, err := s.Insert(context.Background(), g.T)
//		if err != nil {
//			panic(err)
//		}
//	} else {
//		err := s.Update(context.Background(), g.ID, g.T)
//		if err != nil {
//			panic(err)
//		}
//	}
//}
