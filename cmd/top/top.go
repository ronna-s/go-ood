package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/ronna-s/go-ood/pkg/heap"
	"github.com/ronna-s/go-ood/pkg/namegen"
)

type (
	// Artist ...
	Artist struct {
		Name    string
		Listens int
	}
	// Song ...
	Song struct {
		Name    string
		Listens int
	}
)

// Less ...
func (b1 Artist) Less(b2 Artist) bool {
	return b1.Listens > b2.Listens
}

// Less ...
func (p1 Song) Less(p2 Song) bool {
	return p1.Listens > p2.Listens
}

func main() {
	var (
		artists []Artist
		songs   []Song
	)
	for i := 0; i < rand.Intn(1000)+1000; i++ {
		artists = append(artists, Artist{Name: namegen.Generate(), Listens: rand.Intn(851202)})
	}

	for i := 0; i < rand.Intn(1000)+1000; i++ {
		songs = append(songs, Song{Name: namegen.Generate(), Listens: rand.Intn(800917)})
	}

	hartists := heap.New(artists)
	hsongs := heap.New(songs)
	fmt.Println(withColor(cyan, "Our top 10 artists are:"))
	fmt.Println(withColor(cyan, "========================"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %s with %d reads\n", i+1, hartists.Pop().Name, hsongs.Pop().Listens)
	}
	fmt.Println(withColor(purple, "Our top of the pop songs are:"))
	fmt.Println(withColor(purple, "============================="))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %s with %d reads\n", i+1, hsongs.Pop().Name, hsongs.Pop().Listens)
	}
}

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

func withColor(color, s string) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return color + s + "\033[0m"
}
