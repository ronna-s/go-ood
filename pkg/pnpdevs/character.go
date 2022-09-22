package pnpdevs

// Character represents the common fields and abilities of a P&P character
type Character struct {
	X, H int
	Name string
}

// Alive returns false if the Character is dead. The character is dead when its Health is 0
func (c Character) Alive() bool {
	return c.H != 0
}

// GainXP adds XP to the character
func (c *Character) GainXP(xp int) {
	c.X += xp
}

// GainHealth takes in the amount of health to be gained and applies up to a maximum health and down to 0
// of 100. It returns the amount of health actually gained
func (c *Character) GainHealth(health int) int {
	if health+c.H > 100 {
		health = 100 - c.H
	}
	if health+c.H < 0 {
		health = -c.H
	}
	c.H += health
	return health
}

// String ...

// Health ...
func (c Character) Health() int {
	return c.H
}

// XP ...
func (c Character) XP() int {
	return c.X
}
