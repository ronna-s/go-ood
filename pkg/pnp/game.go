package pnp

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	// Game represents a Programmers & Platforms game
	// The purpose of the players is to keep production calm together
	Game struct {
		Players []Player
		Prod    Production
	}

	// Player repsresents a P&P player
	Player interface {
		Alive() bool
		GainXP(int)
		GainHealth(int)
		Skills() []Skill
		Health() int
		XP() int
	}

	// Action ...
	Action = Skill
	// Status represents a P&P production state
	Status interface {
		React(Action) (int, int, Status)
	}
	Production struct {
		Status Status
	}
)

// NewProduction ...
func NewProduction() Production {
	return Production{Calm{}}
}

// React returns the X and health gained by Production's reaction to the player's action
func (p *Production) React(a Action) (xp int, health int) {
	xp, health, p.Status = p.Status.React(a)
	return
}

// Run ...
func (g Game) Run() {
	rand.Seed(time.Now().Unix())
	band := g.Players
	for len(band) != 0 {
		player := band[0]
		band = band[1:]
		skills := player.Skills()
		fmt.Printf("It's %s's turn. %s has %d Health and %d XP. Production's status is '%s'.\n", player, player, player.Health(), player.XP(), g.Prod.Status)
		var choice Skill
		for {
			fmt.Println("Please choose the number of the skill you would like to use:")
			for i := range skills {
				fmt.Printf("%d: %s\n", i+1, skills[i])
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
		if player.Health()+health > 100 {
			health = 100 - player.Health()
		}
		if health >= 0 {
			fmt.Printf("Production liked %s's move. Production's state is now `%s`. Gained: %d XP, %d health\n", player, g.Prod.Status, xp, health)
		} else {
			fmt.Printf("Production DID NOT like %s's move. Production's state is now `%s`. Gained: %d XP, Lost: %d Health\n", player, g.Prod.Status, xp, -health)
		}
		player.GainHealth(health)
		player.GainXP(xp)
		if player.Alive() {
			band = append(band, player)
		} else {
			fmt.Println(fmt.Sprintf("it's so sad that %s is now dead", player))
		}
	}
}
