package main

import (
	"context"
	_ "embed"
	"html/template"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/ronnas/go-ood/pkg/maze"
	"github.com/ronnas/go-ood/pkg/maze/travel"
	"github.com/ronnas/go-ood/pkg/robot"
)

//go:embed resources/maze.tmpl
var tmpl []byte

func main() {
	rand.Seed(time.Now().Unix())
	m := maze.New(rand.Intn(10)+1, rand.Intn(10)+1)
	//fmt.Println(fmt.Sprintf("%#v", m))
	g := robot.New(travel.New(m))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		SolveMaze(&g)
		cancel()
	}()
	<-ctx.Done()
	drawHTML(g, os.Stdout)
}

//go:generate mockery --case=underscore --name=Gopher

// Gopher is an interface to a thing that can move around a maze
type Gopher interface {
	Finished() bool // Has the Gopher reached the target cell?
	Move() error    // The Gopher moves one step in its current direction
	TurnLeft()      // The Gopher will turn left
	TurnRight()     // The Gopher will turn right
}

// SolveMaze is where your code goes to solve our maze
// It takes in g Gopher that knows how to travel.
// See the Gopher interface methods for more details
func SolveMaze(g Gopher) {
	for !g.Finished() {
		g.TurnRight() //try right
		if err := g.Move(); err != nil {
			g.TurnLeft() //try forward
			if err := g.Move(); err != nil {
				g.TurnLeft() //try left
			}
		}
	}
}

// Result represnts the Result of a Maze run
type Result struct {
	maze.Maze
	Steps []robot.Step
}

//drawHTML draws
func drawHTML(g robot.Robot, w io.Writer) {
	res := Result{
		Maze:  g.Maze,
		Steps: g.Steps(),
	}
	const maxMoves = 10000
	if len(res.Steps) > maxMoves {
		res.Steps = res.Steps[:maxMoves]
	}

	t, err := template.New("main").Funcs(template.FuncMap{
		"Upto": func(count int) (items []int) {
			for i := 0; i < count; i++ {
				items = append(items, i)
			}
			return items
		},
		"Incr": func(i int) int {
			return i + 1
		},
		"Passages": func(i int) (s string) {
			c := res.CoordsFromCell(i)
			if res.PathRight(i) {
				s += "right "
			} else if c[maze.X] != res.DimX-1 {
				s += "no-right "
			}
			if res.PathDown(i) {
				s += "down"
			} else if c[maze.Y] != res.DimY-1 {
				s += "no-down"
			}
			return
		},
		"PathDown": func(i int) bool {
			return res.PathDown(i)
		},
	}).Parse(string(tmpl))
	if err != nil {
		panic(err)
	}

	if err := t.ExecuteTemplate(w, "T", res); err != nil {
		panic(err)
	}
}
