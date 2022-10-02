// // Package pnp provides a Platforms and Programmers™ game implementation
package simple

//
//import (
//	"bufio"
//	_ "embed"
//	"fmt"
//	"math/rand"
//	"os"
//	"runtime"
//	"time"
//
//	"github.com/gdamore/tcell/v2"
//	"github.com/rivo/tview"
//)
//
//type (
//	GameComponent struct{}
//	MenuComponent struct{}
//)
//
//func (m *MenuComponent) Render(g *Game) {
//	players := g.Players
//	i := g.Rounds % len(players)
//	prod := g.Prod
//
//	var options = tview.NewList().ShowSecondaryText(false)
//	for i, s := range players[i].Skills() {
//		options.AddItem(s.String(), "", rune(49+i), nil)
//	}
//	choice = 0
//	options.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
//		skill := players[i].Skills()[choice]
//		xp, health := prod.Do(skill)
//		health = players[i].ApplyHealthDiff(health)
//		players[i].ApplyXPDiff(xp)
//		m := tview.NewModal()
//		if health >= 0 {
//			m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", players[i], skill, prod.State, xp, health))
//		} else {
//			m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", players[i], skill, prod.State, xp, health)).SetBackgroundColor(tcell.ColorRed)
//		}
//		m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "ok" {
//				g.Next()
//				i = (i + 1) % len(players)
//				//menuPane = renderMenu(players, i, prod, numTurns)
//				playersPane = renderPlayers(players, i)
//				pages.RemovePage("page")
//				//flex.SetInputCapture(fn)
//				pages.AddPage("page", flex, true, true)
//				pages.SwitchToPage("page")
//			}
//		})
//		pages = pages.AddPage("modal", m, true, true)
//		//pages.SwitchToPage("modal")
//
//	})
//
//	menu := tview.NewFlex().AddItem(options, 0, 1, true)
//
//	menu.SetBorder(true).SetTitle("Select Skill")
//	menuPane = menu
//}
//
//func (g *GameComponent) Render(game *Game) {
//	flex = tview.NewFlex().
//		AddItem(playersPane, 0, 2, false).
//		AddItem(menuPane, 0, 1, true)
//}
//
//type Engine interface {
//	DoTurn(player Player) Production
//	Run()
//}
//
//type (
//	Renderer interface {
//		Render(g Game)
//	}
//	// Game represents a Programmers & Platforms game
//	// The purpose of the players is to keep production calm together
//	Game struct {
//		Engine
//		Name      string
//		Band      []Player
//		Players   []Player
//		Prod      Production
//		Rounds    int
//		Renderers []Renderer
//	}
//
//	// Player represents a P&P player
//	Player interface {
//		Alive() bool
//		ApplyXPDiff(int)
//		ApplyHealthDiff(int) int
//		Skills() []Skill
//		Health() int
//		Art() string
//		XP() int
//	}
//
//	// Action ...
//	Action = Skill
//)
//
//var (
//	app          = tview.NewApplication()
//	pages        = tview.NewPages()
//	menuPane     = tview.NewFlex()
//	playersPane  = tview.NewFlex()
//	welcomeModal = tview.NewModal()
//	flex         = tview.NewFlex()
//	modal        = tview.NewModal()
//	choice       int
//
//	fn func(event *tcell.EventKey) *tcell.EventKey = nil
//)
//
//func (g Game) Render() {
//	for _, r := range g.Renderers {
//		r.Render(g)
//	}
//}
//
////type Page struct {
////	tview.Primitive
////	Game Game
////}
////type Pages struct {
////}
////
////func (p Pages) Start() {
////
////}
////
////func (g Game) Run() {
////	pages := Pages{
////		NewWelcomePage(g),
////		welcomeBandPage(g.Band), //not necessary
////		turnPage,
////		outcomePage,
////	}
////	pages.Start()
////
////	userName := g.Welcome()
////
////	for !g.Finished() {
////		p := g.PopPlayer()
////		g.DoTurn(p)
////		if p.Alive() {
////			g.PushPlayer(p)
////		}
////
////	}
////}
////
////func (g Game) DoTurn(p *Player) {
////	g.Engine.DoTurn(p)
////
////}
//
//var (
//	colorDefault1 = tcell.ColorDarkCyan
//	colorDefault2 = tcell.ColorYellow
//	colorDefault3 = tcell.ColorBlack
//)
//
//// NewGame returns a new P&P game
//func NewGame(tview Engine, name string, prod Production, players ...Player) Game {
//	return Game{Engine: tview, Name: name, Prod: prod, Players: players, Rounds: 0}
//}
//
//func renderMenu(players []Player, i int, prod Production, numTurns int) *tview.Flex {
//	var options = tview.NewList().ShowSecondaryText(false)
//
//	for i, s := range players[i].Skills() {
//		options.AddItem(s.String(), "", rune(49+i), nil)
//	}
//	choice = 0
//	options.SetSelectedFunc(func(choice int, s string, s2 string, r rune) {
//		skill := players[i].Skills()[choice]
//		xp, health := prod.Do(skill)
//		health = players[i].ApplyHealthDiff(health)
//		players[i].ApplyXPDiff(xp)
//		m := tview.NewModal()
//		if health >= 0 {
//			m.SetText(fmt.Sprintf("Production liked %s's move `%s`. Production's state is now `%s`. Gained: %d XP, %d health", players[i], skill, prod.State, xp, health))
//		} else {
//			m.SetText(fmt.Sprintf("Production DID NOT like %s's move `%s`. Production's state is now `%s`. Gained: %d XP, Lost: %d Health", players[i], skill, prod.State, xp, health)).SetBackgroundColor(tcell.ColorRed)
//		}
//		m.AddButtons([]string{"ok"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			if buttonLabel == "ok" {
//				i = (i + 1) % len(players)
//				numTurns++
//				menuPane = renderMenu(players, i, prod, numTurns)
//				playersPane = renderPlayers(players, i)
//				pages.RemovePage("page")
//				//flex.SetInputCapture(fn)
//				pages.AddPage("page", flex, true, true)
//				pages.SwitchToPage("page")
//			}
//		})
//		pages = pages.AddPage("modal", m, true, true)
//		//pages.SwitchToPage("modal")
//
//	})
//
//	menu := tview.NewFlex().AddItem(options, 0, 1, true)
//
//	menu.SetBorder(true).SetTitle("Select Skill")
//	return menu
//}
//
//func renderPlayers(players []Player, i int) *tview.Flex {
//	playersFlex := tview.NewFlex().SetDirection(tview.FlexRow)
//	for j, p := range players {
//		var color tcell.Color
//		if p.Health() > 70 {
//			color = tcell.ColorDarkCyan
//		} else if p.Health() > 50 {
//			color = tcell.ColorYellow
//		} else if p.Health() > 30 {
//			color = tcell.ColorOrange
//		} else {
//			color = tcell.ColorRed
//		}
//		art := tview.NewTextView().SetText(p.Art())
//		art.SetTextColor(color).SetBorder(true).SetBorderPadding(0, 0, 1, 0)
//		if i == j {
//			art.SetTitle(fmt.Sprintf("It's %s's turn", p))
//			art.SetBorderColor(tcell.ColorYellow)
//		}
//
//		playersFlex.AddItem(art, 0, 1, false)
//	}
//	return playersFlex
//}
//
////func renderGame() *tview.Flex {
////	flex := tview.NewFlex().
////		AddItem(playersPane, 0, 2, false).
////		AddItem(menuPane, 0, 1, true)
////
////	return flex
////}
//
//func loadWelcomeForm() *tview.Flex {
//	var bandName string
//
//	newGameText := tview.NewTextView()
//	newGameText.SetText("A band of developers will attempt to survive against PRODUCTION!")
//	gameArt := tview.NewTextView()
//	gameArt.SetText(gamestarted).SetTextColor(tcell.ColorAqua)
//	nameInput := tview.NewInputField().SetLabel("What is the name of your band?  ").SetText("Cool Band").SetFieldTextColor(tcell.ColorBlack).SetFieldBackgroundColor(tcell.ColorDarkCyan).SetFieldWidth(32)
//	nameInput.SetDoneFunc(func(key tcell.Key) {
//		if key != tcell.KeyEnter {
//			return
//		}
//		bandName = nameInput.GetText()
//		welcomeModal.SetText("Hello, " + bandName + "! Are you ready?").SetBackgroundColor(tcell.ColorBlack)
//		welcomeModal.SetTextColor(tcell.ColorDarkCyan)
//		welcomeModal.AddButtons([]string{"Let's do this!"}).SetDoneFunc(func(buttonIndex int, buttonLabel string) {
//			//pages.RemovePage("welcome modal")
//			//pages.RemovePage("load")
//			pages.SwitchToPage("game")
//		})
//		pages.AddPage("welcome modal", welcomeModal, true, true)
//		pages.SwitchToPage("welcome modal")
//	})
//
//	form := tview.NewFlex().
//		SetDirection(tview.FlexRow).
//		AddItem(gameArt, 20, 20, false).
//		AddItem(newGameText, 1, 1, false).
//		AddItem(nameInput, 1, 1, true)
//
//	form.SetBorderPadding(0, 0, 20, 0)
//	form.SetBorder(true).SetTitle("New game started!").SetTitleAlign(tview.AlignLeft)
//	return form
//}
//
//func GameView(i int, players []Player, prod Production, numTurns int) []Player {
//	//menuPane = renderMenu(players, 0, prod, numTurns)
//	playersPane = renderPlayers(players, i)
//	//flex = renderGame()
//	pages.AddPage("load", tview.NewFlex().AddItem(loadWelcomeForm(), 0, 1, true), true, true)
//
//	pages.AddPage("game", flex, true, false)
//	if err := app.SetRoot(pages, true).SetFocus(pages).EnableMouse(true).Run(); err != nil {
//		panic(err)
//	}
//	return nil
//}
//
//func Run(players ...Player) {
//	withColor = func(color, s string) string {
//		return s
//	}
//	GameView(0, players, NewProduction(), 0)
//
//}
//
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
