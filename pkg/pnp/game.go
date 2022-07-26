package pnp

import (
	_ "embed"
	"fmt"
	"io"
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
		Image() string
		XP() int
	}

	// Action ...
	Action = Skill
	// State represents a P&P production state
	State interface {
		React(Action) (int, int, State)
	}
	// Production ...
	Production struct {
		State State
	}
)

// NewGame returns a new P&P game
func NewGame(name string, prod Production, players ...Player) Game {
	return Game{Name: name, Prod: prod, Players: players}
}

// NewProduction ...
func NewProduction() Production {
	return Production{Calm{}}
}

// React returns the X and health gained by Production's reaction to the player's action
func (p *Production) React(a Action) (xp int, health int) {
	xp, health, p.State = p.State.React(a)
	return
}
func clearScr() {
	fmt.Print("\033[H\033[2J")
}

// Run ...
func (g *Game) Run() {
	clearScr()
	rand.Seed(time.Now().Unix())
	var band []Player
	// When loading a new game avoid loading band members who are dead
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

		printImage := func() { fmt.Println(player.Image()) }
		if player.Health() > 70 {
			withColor(green, os.Stdout, printImage)
		} else if player.Health() > 30 {
			withColor(yellow, os.Stdout, printImage)
		} else {
			withColor(red, os.Stdout, printImage)
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
			clearScr()
			withColor(blue, os.Stdout, func() { fmt.Println(gravestone) })
			fmt.Printf("it's so sad that %s is now dead\n", player)
		}
		//fmt.Println("Press enter to continue. [Q] to quit...")
		//b, _ := bufio.NewReader(os.Stdin).ReadByte()
		//if b == 'Q' {
		//	return
		//}
		clearScr()
	}
}

//go:embed resources/gravestone.txt
var gravestone string

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

func withColor(c string, w io.Writer, f func()) {
	if runtime.GOOS == "windows" {
		f()
		return
	}
	defer w.Write([]byte("\033[0m"))
	w.Write([]byte(c))
	f()
}
