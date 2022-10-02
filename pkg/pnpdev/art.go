package pnpdev

import _ "embed" //package provides its on ascii art

//go:embed resources/gopher.txt
var gopherArt string

//go:embed resources/rubyist.txt
var rubyistArt string

//go:embed resources/minion.txt
var minionArt string
