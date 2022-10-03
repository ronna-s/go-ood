// Package engine provides a simple engine for a P&P game
package engine

import (
	"bufio"
	"fmt"
	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnp/engine"
	"os"
	"runtime"
)

type Engine struct {
}

func (e Engine) Reaction(xp int, health int, player pnp.Player, state pnp.State, action pnp.Action, fn func()) {
	defer fn()
	if health >= 0 {
		fmt.Printf("Production liked %s's move. Production's state is now `%s`. Gained: %d XP, %d health\n", player, state, xp, health)
	} else {
		fmt.Printf("Production DID NOT like %s's move. Production's state is now `%s`. Gained: %d XP, Lost: %d Health\n", player, state, xp, -health)
	}
	if !player.Alive() {
		fmt.Println(withColor(purple, engine.Gravestone))
		fmt.Printf("it's so sad that %s is now dead\n", player)
	}
	fmt.Println()
	enter()
	clearScr()
}

func (e Engine) Start() {

}

func (e Engine) RenderGame(players []pnp.Player, p pnp.Player) {
	clearScr()
}

func (e Engine) SelectAction(player pnp.Player, state pnp.State, onSelect func(action pnp.Action)) {
	skills := player.Skills()
	fmt.Printf("It's %s's turn. Production's status is '%s'.\n\n", player, state)

	if player.Health() > 70 {
		fmt.Println(withColor(green, player.Art()))
	} else if player.Health() > 30 {
		fmt.Println(withColor(yellow, player.Art()))
	} else {
		fmt.Println(withColor(red, player.Art()))
	}

	fmt.Println()

	var choice pnp.Action
	for {
		fmt.Println("Please choose the number of the skill you would like to use:")
		for i := range skills {
			fmt.Printf("[%d] %s\n", i+1, skills[i])
		}
		var i int
		if _, err := fmt.Scanln(&i); err != nil {
			fmt.Printf("failed parsing input %s\n", err)
		} else if i < 1 || i > len(skills) {
			fmt.Printf("invalid option %d\n", i)
		} else {
			choice = skills[i-1]
			break
		}
	}
	onSelect(choice)
}

func (e Engine) GameOver() {
	fmt.Println(withColor(cyan, engine.GameOver))
	enter()
}

func (e Engine) GameWon() {
	fmt.Println(withColor(cyan, engine.GameWon))
	enter()
}

func (e Engine) PizzaDelivery(fn func()) {
	defer fn()
	fmt.Println(withColor(cyan, "PIZZA DELIVERY!"))
	fmt.Println(engine.Pizza)
	enter()
}

func (e Engine) Welcome(fn func(string)) {
	fmt.Println(withColor(cyan, engine.Gamestarted))
	fmt.Println("New game started. A band of developers will attempt to survive against Production!")
	fmt.Println("What is the name of your band?")
	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic("error reading band name")
	}
	fmt.Println("Hello, " + string(l) + "! are you ready?")
	enter()
	fn(string(l))
}

var _ pnp.Engine = Engine{}

//// Run ...
//func Run(players ...Player) {
//	fmt.Println(withColor(cyan, gamestarted))
//	fmt.Println("New game started. A band of developers will attempt to survive against Production!")
//	fmt.Println("What is the name of your band?")
//	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
//	if err != nil {
//		panic("error reading band name")
//	}
//
//	g := NewGame(nil, string(l), NewProduction(), players...)
//	clearScr()
//	rand.Seed(time.Now().Unix())
//	band := g.Players
//	for len(band) != 0 {
//		g.Rounds++
//
//		if _, ok := g.Prod.State.(Calm); ok && g.Rounds > 30 {
//			fmt.Println(withColor(cyan, "A notorious business mongrel buys out your company for $50 Billion!"))
//			fmt.Println(withColor(green, "The whole company retires and move to the bahamas!"))
//			fmt.Println(withColor(yellow, "Well done. The game is over!"))
//			return
//		}
//
//		if rand.Intn(30) == 0 {
//			fmt.Println(withColor(cyan, "PIZZA DELIVERY! \nAll players get a pizza, some rest and a health boost!"))
//			fmt.Println(pizza)
//			for i := range band {
//				band[i].ApplyHealthDiff(100)
//			}
//			pressEnter()
//			continue
//		}
//		player := band[0]
//		band = band[1:]
//		skills := player.Skills()
//		fmt.Printf("It's %s's turn. Production's status is '%s'.\n\n", player, g.Prod.State)
//
//		if player.Health() > 70 {
//			fmt.Println(withColor(green, player.Art()))
//		} else if player.Health() > 30 {
//			fmt.Println(withColor(yellow, player.Art()))
//		} else {
//			fmt.Println(withColor(red, player.Art()))
//		}
//
//		fmt.Println()
//
//		var choice Skill
//		for {
//			fmt.Println("Please choose the number of the skill you would like to use:")
//			for i := range skills {
//				fmt.Printf("[%d] %s\n", i+1, skills[i])
//			}
//			var i int
//			if _, err := fmt.Scanln(&i); err != nil {
//				fmt.Printf("failed parsing input %s\n", err)
//			} else if i < 1 || i > len(skills) {
//				fmt.Printf("invalid option %d\n", i)
//			} else {
//				choice = skills[i-1]
//				break
//			}
//		}
//		xp, health := g.Prod.Do(choice)
//		health = player.ApplyHealthDiff(health)
//		player.ApplyXPDiff(xp)
//		if health >= 0 {
//			fmt.Printf("Production liked %s's move. Production's state is now `%s`. Gained: %d XP, %d health\n", player, g.Prod.State, xp, health)
//		} else {
//			fmt.Printf("Production DID NOT like %s's move. Production's state is now `%s`. Gained: %d XP, Lost: %d Health\n", player, g.Prod.State, xp, -health)
//		}
//		fmt.Println()
//
//		if player.Alive() {
//			band = append(band, player)
//		} else {
//			fmt.Println(withColor(purple, gravestone))
//			fmt.Printf("it's so sad that %s is now dead\n", player)
//		}
//		pressEnter()
//	}
//	fmt.Println(withColor(cyan, gameover))
//}
//
//var (
//	red    = "\033[31m"
//	green  = "\033[32m"
//	yellow = "\033[33m"
//	purple = "\033[35m"
//	cyan   = "\033[36m"
//)
//
//func pressEnter() {
//	fmt.Println("Press enter to continue...")
//	_, _ = bufio.NewReader(os.Stdin).ReadByte()
//	clearScr()
//
//}
//
//var withColor = func(color, s string) string {
//	if runtime.GOOS == "windows" {
//		return s
//	}
//	return color + s + "\033[0m"
//}
//
//func clearScr() {
//	fmt.Print("\033[H\033[2J")
//}
//
////go:embed resources/gravestone.txt
//var gravestone string
//
////go:embed resources/gameover.txt
//var gameover string
//
////go:embed resources/gamestarted.txt
//var gamestarted string
//
////go:embed resources/pizza.txt
//var pizza string
//
//type Label struct {
//	*tview.TextView
//}
//
//func (l Label) GetLabel() string {
//	return l.GetText(false)
//}
//
//func (l Label) SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem {
//	return l
//}
//
//func (l Label) GetFieldWidth() int {
//	return 100
//}
//
//func (l Label) SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem {
//	return l
//}

func withColor(color, s string) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return color + s + "\033[0m"
}

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

func clearScr() {
	fmt.Print("\033[H\033[2J")
}

func enter() {
	fmt.Println("Press enter to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()
}
