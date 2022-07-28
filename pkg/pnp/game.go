// Package pnp provides a Platforms and Programmers™ game implementation
package pnp

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
)

type (
	// Game represents a Programmers & Platforms game
	// The purpose of the players is to keep production calm together
	Game struct {
		Name    string
		Players []Player
		Prod    Production
	}

	// Player represents a P&P player
	Player interface {
		Alive() bool
		GainXP(int)
		GainHealth(int) int
		Skills() []Skill
		Health() int
		Art() string
		XP() int
	}

	// Action ...
	Action = Skill
)

// NewGame returns a new P&P game
func NewGame(name string, prod Production, players ...Player) Game {
	return Game{Name: name, Prod: prod, Players: players}
}

// Run ...
func Run() {
	fmt.Println(withColor(cyan, gamestarted))
	fmt.Println("New game started. A band of developers will attempt to survive against Production!")
	fmt.Println("What is the name of your band?")
	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic("error reading band name")
	}

	g := NewGame(string(l), NewProduction(), NewRubyist(), NewGopher())
	clearScr()
	rand.Seed(time.Now().Unix())
	var band []Player
	// When loading a new game avoid loading band members who are already dead
	for _, p := range g.Players {
		if p.Alive() {
			band = append(band, p)
		}
	}
	for len(band) != 0 {
		player := band[0]
		band = band[1:]
		skills := player.Skills()
		fmt.Printf("It's %s's turn. Production's status is '%s'.\n\n", player, g.Prod.State)

		if player.Health() > 70 {
			fmt.Println(withColor(green, player.Art()))
		} else if player.Health() > 30 {
			fmt.Println(withColor(yellow, player.Art()))
		} else {
			fmt.Println(withColor(red, player.Art()))
		}

		fmt.Println()

		var choice Skill
		for {
			fmt.Println("Please choose the number of the skill you would like to use:")
			for i := range skills {
				fmt.Printf("[%d] %s\n", i+1, skills[i])
			}
			var i int
			if _, err := fmt.Scanln(&i); err != nil {
				fmt.Printf("failed parsing input %s\n", err)
			} else if i < 1 || i > len(skills) {
				fmt.Printf("invalid option %d\n", i)
			} else {
				choice = skills[i-1]
				break
			}
		}
		xp, health := g.Prod.React(choice)
		health = player.GainHealth(health)
		player.GainXP(xp)
		if health >= 0 {
			fmt.Printf("Production liked %s's move. Production's state is now `%s`. Gained: %d XP, %d health\n", player, g.Prod.State, xp, health)
		} else {
			fmt.Printf("Production DID NOT like %s's move. Production's state is now `%s`. Gained: %d XP, Lost: %d Health\n", player, g.Prod.State, xp, -health)
		}
		fmt.Println()

		if player.Alive() {
			band = append(band, player)
		} else {
			fmt.Println(withColor(purple, gravestone))
			fmt.Printf("it's so sad that %s is now dead\n", player)
		}
		fmt.Println("Press enter to continue...")
		b, _ := bufio.NewReader(os.Stdin).ReadByte()
		if b == 'Q' {
			return
		}
		clearScr()
	}
	fmt.Println(withColor(cyan, gameover))
}

//go:embed resources/gravestone.txt
var gravestone string

//go:embed resources/gameover.txt
var gameover string

//go:embed resources/gamestarted.txt
var gamestarted string

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

func withColor(color, s string) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return color + s + "\033[0m"
}
func clearScr() {
	fmt.Print("\033[H\033[2J")
}
