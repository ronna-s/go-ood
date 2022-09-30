package pnp

import _ "embed"

type (
	Game struct {
		Players []Player
		Round   int
		Prod    Production
		Alive   []Player
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

	Engine interface {
		Reaction(xp int, health int, player Player, state State, action Action, fn func())
		Start()
		RenderGame(players []Player, p Player)
		SelectAction(p Player, onSelect func(action Action))
		GameOver()
		GameWon()
	}
	// Action ...
	Action = Skill
)

func New(players ...Player) *Game {
	g := Game{Players: players, Round: 0, Prod: NewProduction()}
	g.Alive = make([]Player, len(g.Players))
	copy(g.Alive, g.Players)
	return &g
}

func (g *Game) Run(e Engine) {
	g.DoRound(e)
	e.Start()
}
func (g *Game) DoRound(e Engine) {
	if g.Round > 30 && g.Prod.State == Calm {
		e.GameWon()
		return
	}
	g.Round++
	p := g.PopPlayer()
	if p == nil {
		e.GameOver()
		return
	}
	e.RenderGame(g.Players, p)
	e.SelectAction(p, func(action Action) {
		xp, health := g.React(p, action)
		e.Reaction(xp, health, p, g.Prod.State, action, func() {
			if p.Alive() {
				g.PushPlayer(p)
			}
			g.DoRound(e)
		})
	})
}
func (g *Game) PopPlayer() Player {
	var p Player
	if len(g.Alive) == 0 {
		return nil
	}
	p = g.Alive[0]
	g.Alive = g.Alive[1:]
	return p
}

func (g *Game) PushPlayer(p Player) {
	if p.Alive() {
		g.Alive = append(g.Alive, p)
	}
}

func (g *Game) React(p Player, action Action) (int, int) {
	xp, health := g.Prod.React(action)
	health = p.GainHealth(health)
	p.GainXP(xp)
	return xp, health
}
