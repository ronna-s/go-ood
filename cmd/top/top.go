package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"github.com/ronnas/go-ood/pkg/heap"
	"github.com/ronnas/go-ood/pkg/namegen"
)

type (
	Blogger struct {
		Name  string
		Reads int
	}
	Post struct {
		Name  string
		Reads int
	}
)

// Less ... see https://stackoverflow.com/a/70361353/2730407
func (b1 Blogger) Less(b2 Blogger) bool {
	return b1.Reads > b2.Reads
}

// Less ... see https://stackoverflow.com/a/70361353/2730407
func (p1 Post) Less(p2 Post) bool {
	return p1.Reads > p2.Reads
}

func main() {
	var (
		bloggers []Blogger
		posts    []Post
	)
	for i := 0; i < rand.Intn(1000)+1000; i++ {
		bloggers = append(bloggers, Blogger{Name: namegen.Generate(), Reads: rand.Intn(851202)})
	}

	for i := 0; i < rand.Intn(1000)+1000; i++ {
		posts = append(posts, Post{Name: namegen.Generate(), Reads: rand.Intn(800917)})
	}

	hblogger := heap.New(bloggers)
	hposts := heap.New(posts)
	fmt.Println(withColor(cyan, "Our top 10 bloggers are:"))
	fmt.Println(withColor(cyan, "========================"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %s with %d reads\n", i+1, hblogger.Pop().Name, hposts.Pop().Reads)
	}
	fmt.Println(withColor(purple, "Our top 10 posts are:"))
	fmt.Println(withColor(purple, "====================="))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %s with %d reads\n", i+1, hposts.Pop().Name, hposts.Pop().Reads)
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
func clearScr() {
	fmt.Print("\033[H\033[2J")
}
