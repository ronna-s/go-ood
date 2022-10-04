package pnp

import _ "embed"

type (
	// Game represents a Platforms and Programmers Game
	// where a band of player will attempt to take on production
	Game struct {
		Players  []Player
		Round    int
		Prod     Production
		Alive    []Player
		BandName string
	}

	// Player represents a P&P player
	Player interface {
		// Alive checks if the player is alive
		Alive() bool
		// ApplyXPDiff applies the XP diff to the player's XP
		ApplyXPDiff(xp int) (actual int)
		// ApplyHealthDiff takes in the amount of health to be gained and applies it to the player's health
		// up to a maximum health of 100 and down to 0. It returns the actual difference that was applied
		ApplyHealthDiff(health int) (actual int)
		// Skills returns the skills of the player
		Skills() []Skill
		// Health returns the player's current health level
		Health() int
		// AsciiArt returns the player's ascii art
		AsciiArt() string
		// XP returns the player's XP level
		XP() int
	}

	// Engine represents the game's user interface rendering engine
	Engine interface {
		Reaction(xp int, health int, player Player, state State, action Action, fn func())
		Start()
		RenderGame(players []Player, p Player)
		SelectAction(player Player, state State, onSelect func(action Action))
		GameOver()
		GameWon()
		PizzaDelivery(fn func())
		Welcome(fn func(string))
	}
	// Action ...
	Action = Skill
)

// New returns a new P&P game
func New(players ...Player) *Game {
	g := Game{Players: players, Round: 0, Prod: NewProduction()}
	g.Alive = make([]Player, len(g.Players))
	copy(g.Alive, g.Players)
	return &g
}

// Run starts a new game
func (g *Game) Run(e Engine) {
	g.Welcome(e, func() {
		g.MainLoop(e)
	})
	e.Start()
}
func (g *Game) Welcome(e Engine, fn func()) {
	e.Welcome(func(bandName string) {
		g.BandName = bandName
		g.MainLoop(e)
	})
}

// MainLoop kicks off the next players round
func (g *Game) MainLoop(e Engine) {
	g.Round++
	if g.Round > 1 && Rand(20) == 0 {
		for _, p := range g.Alive {
			p.ApplyHealthDiff(100)
		}
		e.PizzaDelivery(func() {
			g.MainLoop(e)
		})
		return
	}
	if g.Round > 30 && g.Prod.State == Calm {
		e.GameWon()
		return
	}
	p := g.PopPlayer()
	if p == nil {
		e.GameOver()
		return
	}
	e.RenderGame(g.Players, p)
	e.SelectAction(p, g.Prod.State, func(action Action) {
		xp, health := g.React(p, action)
		e.Reaction(xp, health, p, g.Prod.State, action, func() {
			if p.Alive() {
				g.PushPlayer(p)
			}
			g.MainLoop(e)
		})
	})
}

// PopPlayer pops the next player from the queue
func (g *Game) PopPlayer() Player {
	var p Player
	if len(g.Alive) == 0 {
		return nil
	}
	p = g.Alive[0]
	g.Alive = g.Alive[1:]
	return p
}

// PushPlayer returns a player to the end of the queue
func (g *Game) PushPlayer(p Player) {
	if p.Alive() {
		g.Alive = append(g.Alive, p)
	}
}

// React applies a player's action to production,
// applies its outcome (xp, health) to the player and returns it.
func (g *Game) React(p Player, action Action) (int, int) {
	xp, health := g.Prod.React(action)
	health = p.ApplyHealthDiff(health)
	p.ApplyXPDiff(xp)
	return xp, health
}
