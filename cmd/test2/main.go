package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image/jpeg"
	"math"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/harmonica"
	"github.com/charmbracelet/lipgloss"
)

const (
	fps          = 60
	spriteWidth  = 12
	spriteHeight = 5
	frequency    = 7.0
	damping      = 0.15
)

var (
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "246", Dark: "241"}).
		MarginTop(1).
		MarginLeft(2)

	spriteStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#575BD8"))
)

type frameMsg time.Time

func animate() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
		return frameMsg(t)
	})
}

func wait(d time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(d)
		return nil
	}
}

type model struct {
	x      float64
	xVel   float64
	spring harmonica.Spring
}

func (_ model) Init() tea.Cmd {
	return tea.Sequentially(wait(time.Second/2), animate())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	// Step foreward one frame
	case frameMsg:
		const targetX = 60

		// Update x position (and velocity) with our spring.
		m.x, m.xVel = m.spring.Update(m.x, m.xVel, targetX)

		// Quit when we're basically at the target position.
		if math.Abs(m.x-targetX) < 0.01 {
			return m, tea.Sequentially(wait(3/4*time.Second), tea.Quit)
		}

		// Request next frame
		return m, animate()

	default:
		return m, nil
	}
}

func (m model) View() string {
	var out strings.Builder
	fmt.Fprint(&out, "\n")

	offset := int(math.Round(m.x))
	if offset < 0 {
		return ""
	}

	img, _ := jpeg.Decode(bytes.NewReader(gopher))
	rect := img.Bounds()
	for y := 0; y < rect.Dy(); y += 30 {
		var spriteRow string
		for x := 0; x < rect.Dx(); x += 10 {
			c := img.At(x, y)
			rInt, gInt, bInt, _ := c.RGBA()
			//ratio := 255.0 / 65535.0
			//r, g, b := float64(rInt)*ratio, float64(gInt)*ratio, float64(bInt)*ratio
			r, g, b := uint8(rInt>>8), uint8(gInt>>8), uint8(bInt>>8)
			hex := fmt.Sprintf("#%02x%02x%02x", r, g, b)
			spriteRow += lipgloss.NewStyle().
				Foreground(lipgloss.Color(hex)).
				Background(lipgloss.Color("#000000")).Render("/")
		}
		row := strings.Repeat(" ", offset) + spriteRow + "\n"
		fmt.Fprint(&out, row)
	}

	//spriteRow := spriteStyle.Render(strings.Repeat("/", spriteWidth))
	//row := strings.Repeat(" ", offset) + spriteRow + "\n"
	//fmt.Fprint(&out, strings.Repeat(row, spriteHeight))

	fmt.Fprint(&out, helpStyle.Render("Press any key to quit"))

	return out.String()
}

func main() {
	m := model{
		spring: harmonica.NewSpring(harmonica.FPS(fps), frequency, damping),
	}

	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

//go:embed gopher.jpg
var gopher []byte
