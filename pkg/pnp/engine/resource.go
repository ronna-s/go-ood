package engine

import _ "embed"

//go:embed resources/gravestone.txt
var Gravestone string

//go:embed resources/gameover.txt
var GameOver string

//go:embed resources/gamestarted.txt
var Gamestarted string

//go:embed resources/pizza.txt
var Pizza string

//go:embed resources/trophy.txt
var GameWon string
