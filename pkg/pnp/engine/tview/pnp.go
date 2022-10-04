// Package engine provides a simple P&P engine
package engine

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"strings"
	"time"
	"unicode"

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
			SetTitle(fmt.Sprintf("Production is `%s`", e.ProdState))
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
	const pageName = "main"
	e.Pages.RemovePage(pageName)
	view := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(e.RenderPlayers(players, p), 0, 2, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(e.Menu, 0, 1, true).
			AddItem(e.Prod, 0, 1, false), 0, 1, true)
	e.Pages.AddAndSwitchToPage(pageName, view, true)
}

func (e Engine) SelectAction(player pnp.Player, state pnp.State, onSelect func(action pnp.Action)) {
	e.Menu.Clear()
	for i, s := range player.Skills() {
		e.Menu.AddItem(fmt.Sprintf("%s (%d%%)", s.String(), state.Chances(s)), "", rune(49+i), nil)
	}
	e.Menu.SetCurrentItem(len(player.Skills()) - 1)
	e.Menu.SetBorder(true).SetTitle("Select skill")
	e.Menu.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
		skill := player.Skills()[choice]
		onSelect(skill)
	})
}

func (e *Engine) Reaction(xp int, health int, player pnp.Player, state pnp.State, action pnp.Action, fn func()) {
	e.ProdState = state
	m := tview.NewModal()
	skillStr := spaceCamelcase(action.String())
	if health >= 0 {
		m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", player, skillStr, state, xp, health)).SetBackgroundColor(tcell.ColorBlue)
	} else if player.Alive() {
		m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", player, skillStr, state, xp, health)).SetBackgroundColor(tcell.ColorDarkRed)
	} else {
		m.SetText(fmt.Sprintf("%s died (was fired) in the battle against Production. RIP %s. We will always treasure your typos and stuff!!", player, player)).SetBackgroundColor(tcell.ColorPurple)
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
			art.SetText(p.AsciiArt())
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
	case pnp.Annoyed:
		color = tcell.ColorYellow
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
	e.Prod.SetTitle(fmt.Sprintf("Production is `%s`", e.ProdState))
	e.Prod.ScrollToBeginning()
}

func (e *Engine) GameWon() {
	m := NewModal().AddButtons("Yay!").
		SetButtonsAlign(tview.AlignCenter).
		SetText(engine.GameWon).
		SetTextColor(tcell.ColorLime).
		SetBorder(true).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			e.App.Stop()
		})
	m.ResizeItem(m.innerFlex, 0, 5)
	m.innerFlex.ResizeItem(m.modalFlex, 0, 5)

	e.Pages.AddPage("game won", m, true, true)
	//e.Pages.AddAndSwitchToPage("", tview.NewTextView().SetText(engine.GameWon).SetTextColor(tcell.ColorLime), true)
}

func (e Engine) GameOver() {
	m := NewModal().AddButtons("Oh well...").
		SetButtonsAlign(tview.AlignCenter).
		SetText(engine.GameOver).
		SetTextColor(tcell.ColorLime).
		SetBorder(true).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			e.App.Stop()
		})
	m.ResizeItem(m.innerFlex, 0, 3)
	m.innerFlex.ResizeItem(m.modalFlex, 0, 3)

	e.Pages.AddPage("game over", m, true, true)

}

func (e Engine) PizzaDelivery(fn func()) {
	const pageName = "pizza"
	m := NewModal().
		SetText(engine.Pizza).
		SetTextAlign(tview.AlignLeft).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			fn()
			e.Pages.RemovePage(pageName)
		}).
		AddButtons("Thanks, Boss!").
		SetButtonsAlign(tview.AlignCenter)
	m.SetBorder(true)
	m.SetBackgroundColor(tcell.ColorBlack).
		SetTextColor(tcell.ColorGreen)
	m.ResizeItem(m.innerFlex, 0, 3)
	m.innerFlex.ResizeItem(m.modalFlex, 0, 3)
	e.Pages.AddPage(pageName, m, true, true)
}

func (e *Engine) Welcome(fn func(bandName string)) {
	const modalName = "welcome modal"
	newGameText := tview.NewTextView()
	newGameText.SetText("A band of developers will attempt to survive against PRODUCTION!")
	gameArt := tview.NewTextView()
	gameArt.SetText(engine.Gamestarted).SetTextColor(tcell.ColorAqua)
	nameInput := tview.NewInputField().SetLabel("What is the name of your band?  ").SetText("Cool Band").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorDarkCyan).SetFieldWidth(32)
	nameInput.SetDoneFunc(func(key tcell.Key) {
		if key != tcell.KeyEnter {
			return
		}
		bandName := nameInput.GetText()
		welcomeModal := tview.NewModal()
		welcomeModal.SetText("Hello, " + bandName + "! Are you ready?").SetBackgroundColor(tcell.ColorBlack)
		welcomeModal.SetTextColor(tcell.ColorDarkCyan)
		welcomeModal.AddButtons([]string{"Let's do this!"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			e.Pages.RemovePage(modalName)
			e.Pages.RemovePage("load")
			fn(bandName)
		})
		e.Pages.AddAndSwitchToPage(modalName, welcomeModal, true)
	})

	form := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(gameArt, 20, 20, false).
		AddItem(newGameText, 1, 1, false).
		AddItem(nameInput, 1, 1, true)

	form.SetBorderPadding(0, 0, 20, 0)
	form.SetBorder(true).SetTitle("New game started!").SetTitleAlign(tview.AlignLeft)
	e.Pages.AddAndSwitchToPage("load", tview.NewFlex().AddItem(form, 0, 1, true), true)
}

type Modal struct {
	*tview.Flex
	text      *tview.TextView
	form      *tview.Form
	innerFlex *tview.Flex
	modalFlex *tview.Flex
	done      func(idx int, label string)
}

func NewModal() *Modal {
	m := &Modal{form: tview.NewForm(), text: tview.NewTextView()}
	m.modalFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(m.text, 0, 4, false).
		AddItem(m.form, 0, 1, true)
	m.innerFlex = tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(m.modalFlex, 0, 2, true).
		AddItem(nil, 0, 1, false)
	m.Flex = tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(m.innerFlex, 0, 2, true).
		AddItem(nil, 0, 1, false)
	return m
}

func (m *Modal) AddButtons(labels ...string) *Modal {
	for i, label := range labels {
		m.form.AddButton(label, func() {
			m.done(i, label)
		})
	}
	return m
}

func (m *Modal) SetText(text string) *Modal {
	m.text.SetText(text)
	return m
}

func (m *Modal) SetTextAlign(align int) *Modal {
	m.text.SetTextAlign(align)
	return m

}
func (m *Modal) SetButtonsAlign(align int) *Modal {
	m.form.SetButtonsAlign(align)
	return m
}

func (m *Modal) SetBackgroundColor(color tcell.Color) *Modal {
	m.modalFlex.SetBackgroundColor(color)
	m.form.SetBackgroundColor(color)
	m.text.SetBackgroundColor(color)
	return m
}

func (m *Modal) SetBorder(show bool) *Modal {
	m.modalFlex.SetBorder(show)
	return m
}

func (m *Modal) SetDoneFunc(done func(buttonIndex int, buttonLabel string)) *Modal {
	m.done = done
	return m
}

func (m *Modal) SetTextColor(color tcell.Color) *Modal {
	m.text.SetTextColor(color)
	return m
}
