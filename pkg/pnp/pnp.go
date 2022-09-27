package pnp

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (

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

const colorDefault = tcell.ColorAqua

var (
	bandName     string
	playersView  *tview.Flex
	playersViews []*PlayerArt
	menuOptions  *tview.List
	app          *tview.Application
	pages        *tview.Pages
	welcomeView  *tview.Modal
	newGameView  tview.Primitive
	mainView     *tview.Flex
	prodView     *tview.TextView
)

func renderProdView(prod *Production) {
	var color tcell.Color
	switch prod.State {
	case Calm:
		color = tcell.ColorGreen
	case SlightlyAnnoyed:
		color = tcell.ColorYellow
	case VeryAnnoyed:
		color = tcell.ColorOrange
	case Enraged:
		color = tcell.ColorRed
	case Legacy:
		color = tcell.ColorPurple
	}

	text := prodView.GetText(false)
	for i := 0; i < 10; i++ {
		c := string(rune(Rand(128-48) + 48))
		r := Rand(len(text))
		text = text[:r] + c + text[r+1:]
	}
	prodView.SetText(text).SetTextColor(color)
	prodView.SetTitle(fmt.Sprintf("Production is `%s`", strings.ToLower(StrToSentence(prod.State.String()))))
}
func onBandNameSelection(nameInput *tview.InputField) func(key tcell.Key) {
	return func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		bandName = nameInput.GetText()
		welcomeView.SetText("Hello, " + bandName + "! Are you ready?").SetBackgroundColor(tcell.ColorBlack)
		welcomeView.SetTextColor(colorDefault)
		welcomeView.AddButtons([]string{"Let's do this!"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.AddAndSwitchToPage("main", mainView, true)
		})
		pages.AddAndSwitchToPage("welcome modal", welcomeView, true)
	}
}

func renderNewGameView() {
	newGameText := tview.NewTextView()
	newGameText.SetText("A band of developers will attempt to survive against PRODUCTION!")
	gameArt := tview.NewTextView()
	gameArt.SetText(gamestarted).SetTextColor(colorDefault)
	nameInput := tview.NewInputField().SetLabel("What is the name of your band?  ").
		SetText("My Cool Band").
		SetFieldTextColor(tcell.ColorBlack).
		SetFieldBackgroundColor(tcell.ColorWhite).
		SetFieldWidth(32)
	nameInput.SetDoneFunc(onBandNameSelection(nameInput))

	form := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gameArt, 20, 20, false).
		AddItem(newGameText, 1, 1, false).
		AddItem(nameInput, 1, 1, true)

	form.SetBorderPadding(0, 0, 20, 0)
	form.SetBorder(true).SetTitle("New game started!").SetTitleAlign(tview.AlignLeft)
	welcomeView = tview.NewModal()
	newGameView = form
}

type PlayerArt struct {
	Player
	*tview.TextView
}

func (art PlayerArt) Render() {
	art.SetBorderColor(tcell.ColorWhite)
	art.SetTitle("")
	if art.Alive() {
		art.SetText(art.Art())
	} else {
		art.SetText(gravestone)
	}
}

func (art PlayerArt) onTurn() {
	art.SetTitle(fmt.Sprintf("It's %s's turn", art.Player))
	art.SetBorderColor(tcell.ColorYellow)

}
func renderMenu(players []Player, current int, prod *Production) {
	menuOptions.Clear()
	for i, s := range players[current].Skills() {
		menuOptions.AddItem(s.String(), "", rune(49+i), nil)
	}
	menuOptions.SetCurrentItem(len(players[current].Skills()) - 1)
	menuOptions.SetBorder(true).SetTitle("Select skill")
	menuOptions.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
		skill := players[current].Skills()[choice]
		xp, health := prod.React(skill)
		health = players[current].GainHealth(health)
		players[current].GainXP(xp)
		m := tview.NewModal()
		state := StrToSentence(prod.State.String())
		if health >= 0 {
			m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", players[current], skill, state, xp, health)).SetBackgroundColor(tcell.ColorBlue)
		} else if players[current].Alive() {
			m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", players[current], skill, state, xp, health)).SetBackgroundColor(tcell.ColorDarkRed)
		} else {
			m.SetText(fmt.Sprintf("%s died in the battle against Production. RIP %s. We will always treasure your typos and stuff!!", players[current], players[current])).SetBackgroundColor(tcell.ColorPurple)
		}
		m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "ok" {
				playersViews[current].Render()
				current = NextLivingPlayer(playersViews, current)
				if current == -1 {
					pages.AddAndSwitchToPage("", tview.NewTextView().SetText(gameover), true)
				} else {
					playersViews[current].onTurn()
					renderMenu(players, current, prod)
					pages.RemovePage("modal")
				}
			}
		})
		pages = pages.AddPage("modal", m, true, true)
	})
}

func NextLivingPlayer(views []*PlayerArt, current int) int {
	for i := current + 1; i < len(views); i++ {
		if views[i].Alive() {
			return i
		}
	}
	for i := 0; i <= current; i++ {
		if views[i].Alive() {
			return i
		}
	}
	return -1
}
func renderPlayers(players []Player) {
	playersView = tview.NewFlex().SetDirection(tview.FlexRow)
	for _, p := range players {
		var color tcell.Color
		if p.Health() > 70 {
			color = colorDefault
		} else if p.Health() > 50 {
			color = tcell.ColorYellow
		} else if p.Health() > 30 {
			color = tcell.ColorOrange
		} else {
			color = tcell.ColorRed
		}
		//todo: refresh this text
		art := &PlayerArt{TextView: tview.NewTextView().SetText(p.Art()), Player: p}
		art.SetChangedFunc(func() {
			app.Draw()
		})

		playersViews = append(playersViews, art)

		art.SetTextColor(color).SetBorder(true).SetBorderPadding(0, 0, 1, 0)
		playersView.AddItem(art, 0, 1, false)
	}
}
func renderMainView(players []Player) {
	var prod Production
	playersView = tview.NewFlex().SetDirection(tview.FlexRow)
	menuOptions = tview.NewList().ShowSecondaryText(false)
	prodView = tview.NewTextView().SetText(strings.Repeat("A", 1000))
	prodView.SetBorder(true)
	rightPane := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(menuOptions, 0, 1, true).
		AddItem(prodView, 0, 1, true)
	renderMenu(players, 0, &prod)
	renderPlayers(players)
	playersViews[0].onTurn()
	renderProdView(&prod)
	prodView.SetChangedFunc(func() {
		app.Draw()
	})
	go func() {
		time.Sleep(time.Second)
		for {
			time.Sleep(time.Millisecond * 10)
			renderProdView(&prod)
		}
	}()
	mainView = tview.NewFlex().
		AddItem(playersView, 0, 2, false).
		AddItem(rightPane, 0, 1, true) // menuItem
	mainView.SetTitle(bandName)
}

func Run(name string, players ...Player) {
	app = tview.NewApplication()
	withColor = func(color, s string) string { return "" }
	renderNewGameView()
	renderMainView(players)
	pages = tview.NewPages()
	pages.AddAndSwitchToPage("new", newGameView, true)

	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

var (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	cyan   = "\033[36m"
)

var withColor = func(color, s string) string {
	if runtime.GOOS == "windows" {
		return s
	}
	return color + s + "\033[0m"
}

//go:embed resources/gravestone.txt
var gravestone string

//go:embed resources/gameover.txt
var gameover string

//go:embed resources/gamestarted.txt
var gamestarted string

//go:embed resources/pizza.txt
var pizza string

func StrToSentence(s string) string {
	var (
		words []string
		curr  string
	)
	for _, r := range s {
		if unicode.IsUpper(r) && curr != "" {
			words = append(words, curr)
			curr = ""
		}
		curr += string(r)
	}
	words = append(words, curr)

	return strings.Join(words, " ")
}
