// Package pnp provides a Platforms and Programmers™ game implementation
package pnp

import (
	"bufio"
	_ "embed"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (
	// Game represents a Programmers & Platforms game
	// The purpose of the players is to keep production calm together
	Game struct {
		Name    string
		Players []Player
		Prod    Production
		Turns   int
	}

	// Player represents a P&P player
	Player interface {
		Alive() bool
		GainXP(int)
		GainHealth(int) int
		Skills() []Skill
		Health() int
		Art() string
		XP() int
	}

	// Action ...
	Action = Skill
)

var (
	app          = tview.NewApplication()
	pages        = tview.NewPages()
	menuPane     = tview.NewFlex()
	playersPane  = tview.NewFlex()
	welcomeModal = tview.NewModal()
	flex         = tview.NewFlex()
	modal        = tview.NewModal()
	choice       int

	fn func(event *tcell.EventKey) *tcell.EventKey = nil
)

// NewGame returns a new P&P game
func NewGame(name string, prod Production, players ...Player) Game {
	return Game{Name: name, Prod: prod, Players: players, Turns: 0}
}

func renderMenu(players []Player, i int, prod Production, numTurns int) *tview.Flex {
	var options = tview.NewList().ShowSecondaryText(false)

	for i, s := range players[i].Skills() {
		options.AddItem(s.String(), "", rune(49+i), nil)
	}
	choice = 0
	options.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
		skill := players[i].Skills()[choice]
		xp, health := prod.React(skill)
		health = players[i].GainHealth(health)
		players[i].GainXP(xp)
		m := tview.NewModal()
		if health >= 0 {
			m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", players[i], skill, prod.State, xp, health))
		} else {
			m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", players[i], skill, prod.State, xp, health)).SetBackgroundColor(tcell.ColorRed)
		}
		m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				i = (i + 1) % len(players)
				numTurns++
				menuPane = renderMenu(players, i, prod, numTurns)
				playersPane = renderPlayers(players, i)
				flex = renderGame()
				pages.RemovePage("page")
				//flex.SetInputCapture(fn)
				pages.AddPage("page", flex, true, true)
				pages.SwitchToPage("page")
			}
		})
		pages = pages.AddPage("modal", m, true, true)
		//pages.SwitchToPage("modal")

	})

	menu := tview.NewFlex().AddItem(options, 0, 1, true)

	menu.SetBorder(true).SetTitle("Select Skill")
	return menu
}

func renderPlayers(players []Player, i int) *tview.Flex {
	playersFlex := tview.NewFlex().SetDirection(tview.FlexRow)
	for j, p := range players {
		var color tcell.Color
		if p.Health() > 70 {
			color = tcell.ColorGreen
		} else if p.Health() > 50 {
			color = tcell.ColorYellow
		} else if p.Health() > 30 {
			color = tcell.ColorOrange
		} else {
			color = tcell.ColorRed
		}
		art := tview.NewTextView().SetText(p.Art())
		art.SetTextColor(color).SetBorder(true)
		if i == j {
			art.SetTitle(fmt.Sprintf("It's %s's turn", p))
			art.SetBorderColor(tcell.ColorPink)
		}

		playersFlex.AddItem(art, 0, 1, false)
	}
	return playersFlex
}

func renderGame() *tview.Flex {
	flex := tview.NewFlex().
		AddItem(playersPane, 0, 2, false).
		AddItem(menuPane, 0, 1, true)

	return flex
}

func loadWelcomeForm(fn func(event *tcell.EventKey) *tcell.EventKey) *tview.Form {
	label1 := &Label{tview.NewTextView()}
	label1.SetText("A band of developers will attempt to survive against PRODUCTION!")
	nameInput := tview.NewInputField().SetLabel("What is the name of your band?").SetText("Cool Band")
	var bandName string

	form := tview.NewForm().
		AddFormItem(label1).
		AddFormItem(nameInput).
		SetButtonsAlign(tview.AlignCenter).
		SetFocus(1).
		AddButton("Continue", func() {
			bandName = nameInput.GetText()
			welcomeModal.SetText("Hello, " + bandName + "!")
			welcomeModal.SetTitle("New game").SetTitleColor(tcell.ColorOrangeRed)
			welcomeModal.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.RemovePage("welcome modal")
				pages.RemovePage("load")
				pages.SwitchToPage("game")
				//flex.SetInputCapture(fn)
			})
			pages.AddPage("welcome modal", welcomeModal, true, true)
			pages.SwitchToPage("welcome modal")
		}).
		AddButton("Quit", func() {
			app.Stop()
		})

	form.SetBorder(true).SetTitle("New game started!").SetTitleAlign(tview.AlignCenter).SetTitleColor(tcell.ColorLime)
	return form
}

func GameView(i int, players []Player, prod Production, numTurns int) []Player {
	menuPane = renderMenu(players, 0, prod, numTurns)
	playersPane = renderPlayers(players, i)
	flex = renderGame()
	fn = func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEnter {
			skill := players[i].Skills()[choice]
			xp, health := prod.React(skill)
			health = players[i].GainHealth(health)
			players[i].GainXP(xp)
			m := tview.NewModal()
			if health >= 0 {
				m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", players[i], skill.String()+strconv.Itoa(choice), prod.State, xp, health))
			} else {
				m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", players[i], skill, prod.State, xp, health)).SetBackgroundColor(tcell.ColorRed)
			}
			m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "ok" {
					i = (i + 1) % len(players)
					numTurns++
					menuPane = renderMenu(players, i, prod, numTurns)
					playersPane = renderPlayers(players, i)
					flex = renderGame()
					//flex.SetInputCapture(fn)
					pages.AddPage("page", flex, true, true)
					pages.SwitchToPage("page")
				}
			})
			pages = pages.AddPage("modal", m, true, true)
			pages.SwitchToPage("modal")
		}
		return event
	}

	pages.AddPage("load", tview.NewFlex().AddItem(loadWelcomeForm(fn), 0, 1, true), true, true)

	pages.AddPage("game", flex, true, false)
	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	return nil
}

func Run2(players ...Player) {
	withColor = func(color, s string) string {
		return s
	}
	GameView(0, players, NewProduction(), 0)

}

// Run ...
func Run(players ...Player) {
	fmt.Println(withColor(cyan, gamestarted))
	fmt.Println("New game started. A band of developers will attempt to survive against Production!")
	fmt.Println("What is the name of your band?")
	l, _, err := bufio.NewReader(os.Stdin).ReadLine()
	if err != nil {
		panic("error reading band name")
	}

	g := NewGame(string(l), NewProduction(), players...)
	clearScr()
	rand.Seed(time.Now().Unix())
	band := g.Players
	for len(band)+1 != 0 {
		g.Turns++

		if _, ok := g.Prod.State.(Calm); ok && g.Turns > 30 {
			fmt.Println(withColor(cyan, "A notorious business mongrel buys out your company for $50 Billion!"))
			fmt.Println(withColor(green, "The whole company retires and move to the bahamas!"))
			fmt.Println(withColor(yellow, "Well done. The game is over!"))
			return
		}

		if rand.Intn(30) == 0 {
			fmt.Println(withColor(cyan, "PIZZA DELIVERY! \nAll players get a pizza, some rest and a health boost!"))
			fmt.Println(pizza)
			for i := range band {
				band[i].GainHealth(100)
			}
			pressEnter()
			continue
		}
		player := band[0]
		band = band[1:]
		skills := player.Skills()
		fmt.Printf("It's %s's turn. Production's status is '%s'.\n\n", player, g.Prod.State)

		if player.Health() > 70 {
			fmt.Println(withColor(green, player.Art()))
		} else if player.Health() > 30 {
			fmt.Println(withColor(yellow, player.Art()))
		} else {
			fmt.Println(withColor(red, player.Art()))
		}

		fmt.Println()

		var choice Skill
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
		xp, health := g.Prod.React(choice)
		health = player.GainHealth(health)
		player.GainXP(xp)
		if health >= 0 {
			fmt.Printf("Production liked %s's move. Production's state is now `%s`. Gained: %d XP, %d health\n", player, g.Prod.State, xp, health)
		} else {
			fmt.Printf("Production DID NOT like %s's move. Production's state is now `%s`. Gained: %d XP, Lost: %d Health\n", player, g.Prod.State, xp, -health)
		}
		fmt.Println()

		if player.Alive() {
			band = append(band, player)
		} else {
			fmt.Println(withColor(purple, gravestone))
			fmt.Printf("it's so sad that %s is now dead\n", player)
		}
		pressEnter()
	}
	fmt.Println(withColor(cyan, gameover))
}

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

func pressEnter() {
	fmt.Println("Press enter to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadByte()
	clearScr()

}

var withColor = func(color, s string) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return color + s + "\033[0m"
}

func clearScr() {
	fmt.Print("\033[H\033[2J")
}

//go:embed resources/gravestone.txt
var gravestone string

//go:embed resources/gameover.txt
var gameover string

//go:embed resources/gamestarted.txt
var gamestarted string

//go:embed resources/pizza.txt
var pizza string

type Label struct {
	*tview.TextView
}

func (l Label) GetLabel() string {
	return l.GetText(false)
}

func (l Label) SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem {
	return l
}

func (l Label) GetFieldWidth() int {
	return 100
}

func (l Label) SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem {
	return l
}
