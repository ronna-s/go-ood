package pnp

import (
	"fmt"
	"sort"
	"strings"
)

type (
	// Gopher represents the P&P role of the Gopher
	Gopher struct {
		Character
	}
)

// NewGopher ...
func NewGopher() *Gopher {
	return &Gopher{Character{H: 100, Name: "Gopher"}}
}

// Skills returns the list of abilities the Gopher has
func (g Gopher) Skills() []Skill {
	abs := []Skill{TypeSafety}
	switch {
	case g.X >= 100:
		abs = append(abs, Generics)
		fallthrough
	case g.X >= 10:
		abs = append(abs, Interface)
	}
	sort.Slice(abs, func(i, j int) bool { return abs[i] < abs[j] })
	return abs
}

func (g Gopher) Image() string {
	return fmt.Sprintf(strings.Join([]string{
		"                                    ",
		"                                    ",
		"           ,_---~~~~~----.          ",
		"    _,,_,*^____      _____``*g*\"*, ",
		"   / __/ /'     ^.  /      \\ ^@q   f               HEALTH=%d",
		"  [  @f | @))    |  | @))   l  0 _/  ",
		"   \\`/   \\~____ / __ \\_____/    \\   ",
		"    |           _l__l_           I  ",
		"    }          [______]           I ",
		"    ]            | | |            |                XP=%d",
		"    ]             ~ ~             | ",
		"    |                            |  ",
		"     |                           |  ",
		"     |                           |  ",
	}, "\n"), g.H, g.X)
}
