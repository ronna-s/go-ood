package engine

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ronna-s/go-ood/pkg/pnp"
	"github.com/ronna-s/go-ood/pkg/pnp/engine"
)

type Engine struct {
	App       *tview.Application
	Pages     *tview.Pages
	Menu      *tview.List
	Prod      *tview.TextView
	ProdState pnp.State
}

func New() *Engine {
	return &Engine{
		App:   tview.NewApplication(),
		Pages: tview.NewPages(),
		Menu:  tview.NewList(),
		Prod:  tview.NewTextView(),
	}
}
func (e *Engine) Start() {
	go func() {
		e.Prod.SetText(strings.Repeat("A", 2000)).
			SetTextColor(tcell.ColorGreen).
			SetBorder(true).
			SetTitle(fmt.Sprintf("Production is `%s`", strings.ToLower(spaceCamelcase(e.ProdState.String()))))
		e.Prod.SetChangedFunc(func() {
			e.App.Draw()
		})
		time.Sleep(time.Second)
		for {
			time.Sleep(time.Millisecond * 10)
			e.RenderProd()
		}
	}()
	if err := e.App.SetRoot(e.Pages, true).SetFocus(e.Pages).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
func (e *Engine) Stop() {
	e.App.Stop()
}

func (e Engine) RenderGame(players []pnp.Player, p pnp.Player) {
	e.Pages.RemovePage("game")
	view := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(e.RenderPlayers(players, p), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(e.Menu, 0, 1, true).
			AddItem(e.Prod, 0, 1, false), 0, 1, true)
	e.Pages.AddAndSwitchToPage("game", view, true)
}

func (e Engine) SelectAction(p pnp.Player, onSelect func(action pnp.Action)) {
	e.Menu.Clear()
	for i, s := range p.Skills() {
		e.Menu.AddItem(s.String(), "", rune(49+i), nil)
	}
	e.Menu.SetCurrentItem(len(p.Skills()) - 1)
	e.Menu.SetBorder(true).SetTitle("Select skill")
	e.Menu.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
		skill := p.Skills()[choice]
		onSelect(skill)
	})
}

func (e Engine) Reaction(xp int, health int, player pnp.Player, state pnp.State, action pnp.Action, fn func()) {
	e.ProdState = state
	m := tview.NewModal()
	stateStr := spaceCamelcase(state.String())
	skillStr := spaceCamelcase(action.String())
	if health >= 0 {
		m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", player, skillStr, stateStr, xp, health)).SetBackgroundColor(tcell.ColorBlue)
	} else if player.Alive() {
		m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", player, skillStr, stateStr, xp, health)).SetBackgroundColor(tcell.ColorDarkRed)
	} else {
		m.SetText(fmt.Sprintf("%s died in the battle against Production. RIP %s. We will always treasure your typos and stuff!!", player, player)).SetBackgroundColor(tcell.ColorPurple)
	}
	m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "ok" {
			e.Pages.RemovePage("modal")
			fn()
		}
	})
	e.Pages.AddPage("modal", m, true, true)
}

func spaceCamelcase(s string) string {
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

func (e *Engine) RenderPlayers(players []pnp.Player, current pnp.Player) *tview.Flex {
	playersView := tview.NewFlex().SetDirection(tview.FlexRow)
	for _, p := range players {
		var color tcell.Color
		if p.Health() > 70 {
			color = tcell.ColorAqua
		} else if p.Health() > 50 {
			color = tcell.ColorYellow
		} else if p.Health() > 30 {
			color = tcell.ColorOrange
		} else {
			color = tcell.ColorRed
		}
		art := tview.NewTextView()
		art.SetBorderColor(tcell.ColorWhite)

		if p.Alive() {
			art.SetText(p.Art())
		} else {
			art.SetText(engine.Gravestone).SetTextColor(tcell.ColorPurple)
		}
		if p == current {
			art.SetTitle(fmt.Sprintf("It's %s's turn", p)).
				SetBorderColor(tcell.ColorYellow)
		}

		art.SetTextColor(color).SetBorder(true).SetBorderPadding(0, 0, 1, 0)
		playersView.AddItem(art, 0, 1, false)
	}
	return playersView
}

func (e *Engine) RenderProd() {
	var color tcell.Color
	switch e.ProdState {
	case pnp.Calm:
		color = tcell.ColorGreen
	case pnp.SlightlyAnnoyed:
		color = tcell.ColorYellow
	case pnp.VeryAnnoyed:
		color = tcell.ColorOrange
	case pnp.Enraged:
		color = tcell.ColorRed
	case pnp.Legacy:
		color = tcell.ColorPurple
	}

	text := e.Prod.GetText(false)
	for i := 0; i < 10; i++ {
		c := string(rune(pnp.Rand(128-48) + 48))
		r := pnp.Rand(len(text))
		text = text[:r] + c + text[r+1:]
	}
	e.Prod.SetText(text).SetTextColor(color)
	e.Prod.SetTitle(fmt.Sprintf("Production is `%s`", strings.ToLower(spaceCamelcase(e.ProdState.String()))))
	e.Prod.ScrollToBeginning()
}

func NewGame() {

	//─=≡Σ((( つ•̀ω•́)つ LET'S GO!
	// d(-_^)
	// ᕦ(òᴥó)ᕥ
	// ᕕ( ᐛ )ᕗ

}

func (e *Engine) GameWon() {
	e.Pages.AddAndSwitchToPage("", tview.NewTextView().SetText(engine.GameWon).SetTextColor(tcell.ColorLime), true)
}

func (e Engine) GameOver() {
	e.Pages.AddAndSwitchToPage("", tview.NewTextView().SetText(engine.GameOver).SetTextColor(tcell.ColorLime), true)
}
