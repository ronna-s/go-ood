package pnpdev

// Character represents the common fields and abilities of a P&P character
type Character struct {
	H, X int
}

// Alive returns false if the Character is dead. The character is dead when its Health is 0
func (c Character) Alive() bool {
	return c.H > 0
}

// ApplyXPDiff adds XP to the character
func (c *Character) ApplyXPDiff(xp int) int {
	c.X += xp
	return xp
}

// ApplyHealthDiff takes in the amount of health to be gained and applies up to a maximum health of 100 and down to 0
// It returns the amount of health actually gained
func (c *Character) ApplyHealthDiff(health int) int {
	sumH := c.H + health
	if sumH > 100 {
		health = 100 - c.H
		c.H = 100
		return health
	}
	if sumH < 0 {
		health = 0 - c.H
		c.H = 0
		return health
	}
	c.H = sumH
	return health
}

// Health ...
func (c Character) Health() int {
	return c.H
}

// XP ...
func (c Character) XP() int {
	return c.X
}
