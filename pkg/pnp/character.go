package pnp

// Character represents the common fields and abilities of a P&P charachter
type Character struct {
	Name string // The name of the character
	X    int    // Experience points collected
	H    int    // A 0-100 value of the characters health percentage
}

// Alive returns false if the Character is dead
func (c Character) Alive() bool {
	return c.H > 0
}

// GainXP ...
func (c *Character) GainXP(xp int) {
	c.X += xp
}

// GainHealth ...
func (c *Character) GainHealth(health int) {
	c.H += health
}

// String ...
func (c Character) String() string {
	return c.Name
}

// Health ...
func (c Character) Health() int {
	return c.H
}

// XP ...
func (c Character) XP() int {
	return c.X
}
