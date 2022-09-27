package pnp

type (
	// Production ...
	Production struct {
		State State
	}
	// State represents a P&P Production state
	State int
)

// NewProduction ...
func NewProduction() Production {
	return Production{}
}

// React returns the X and health gained by Production's reaction to the player's action
func (p *Production) React(a Action) (xp int, health int) {
	xp, health, p.State = p.State.React(a)
	return
}
