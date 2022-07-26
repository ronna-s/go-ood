package pnp

import (
	"fmt"
	"sort"
	"strings"
)

type (
	// Rubyist represents the P&P role of the Rubyist
	Rubyist struct {
		Character
	}
)

// NewRubyist ...
func NewRubyist() *Rubyist {
	return &Rubyist{Character{H: 100, Name: "Rubyist"}}
}
func (r Rubyist) Image() string {
	return fmt.Sprintf(strings.Join([]string{
		"              ~~++++:  .:+++++. ",
		"           ~+++~.. .+=+o:    ~=~",
		"         :++~       .o.~++~   ~o",
		"       :+:          =:   .++~ ~o",
		"     ~=+          .=+......:=o==                   HEALTH=%d",
		"    ~o~          ~=+=+:::::~~~=+",
		"   ~o.         ~++. o~        ++",
		"   o~        ~++.   ~o        o~",
		"   =:    .~+=+.      o~      .o~",
		"   .o++++:~o+:++:~.. ~o      ~o.                   XP=%d",
		"   +o+    :=   .~~:+++o:     ~o ",
		"  ++ =+  .o.         ...     ++ ",
		"  o~  o: =+                  ++ ",
		"  ~=~ .=+=        ......~~~~~o: ",
	}, "\n"), r.H, r.X)
}

// Abilities returns the list of abilities the Rubyist has
func (r Rubyist) Skills() []Skill {
	abs := []Skill{DuckTyping}
	switch {
	case r.X >= 100:
		abs = append(abs, DarkMagic)
		fallthrough
	case r.X >= 10:
		abs = append(abs, Module)
	}
	sort.Slice(abs, func(i, j int) bool { return abs[i] < abs[j] })
	return abs
}
