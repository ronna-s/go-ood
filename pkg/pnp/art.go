package pnp

import _ "embed" //package provides its on ascii art

//go:embed resources/gopher.txt
var gopherArt string

//go:embed resources/rubyist.txt
var rubyistArt string

//go:embed resources/gravestone.txt
var gravestone string

//go:embed resources/gameover.txt
var gameover string

//go:embed resources/gamestarted.txt
var gamestarted string

//go:embed resources/pizza.txt
var pizza string
