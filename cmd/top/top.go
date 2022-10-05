package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	rand.Seed(time.Now().Unix())
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

	hartists := heap.New[Artist](artists)
	hsongs := heap.New[Song](songs)

	artistsCol := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "Artist", Width: 30},
		{Title: "Listens", Width: 10},
	}
	songsCol := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "Song", Width: 30},
		{Title: "Listens", Width: 10},
	}
	var rows []table.Row
	for i := 0; i < 100; i++ {
		artist := hartists.Pop()
		rows = append(rows, []string{strconv.Itoa(i + 1), artist.Name, strconv.Itoa(artist.Listens)})
	}

	artistsTable := table.New(
		table.WithColumns(artistsCol),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	artistsTable.SetStyles(s)

	rows = nil
	for i := 0; i < 100; i++ {
		song := hsongs.Pop()
		rows = append(rows, []string{strconv.Itoa(i + 1), song.Name, strconv.Itoa(song.Listens)})
	}

	songsTable := table.New(
		table.WithColumns(songsCol),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("32")).
		Bold(false)
	songsTable.SetStyles(s)

	m := model{artistsTable, songsTable}
	if err := tea.NewProgram(m).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

var artistsBaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("132"))
var songsBaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("178"))

type model struct {
	artists table.Model
	songs   table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.artists.Focused() {
				m.artists.Blur()
			} else {
				m.artists.Focus()
			}
			if m.songs.Focused() {
				m.songs.Blur()
			} else {
				m.songs.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.artists.Focused() {
				m.artists.Blur()
				m.songs.Focus()
			} else {
				m.songs.Blur()
				m.artists.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.artists.SelectedRow()[1]),
			)
		}
	}
	if m.artists.Focused() {
		m.artists, cmd = m.artists.Update(msg)
	} else if m.songs.Focused() {
		m.songs, cmd = m.songs.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	return artistsBaseStyle.Render(m.artists.View()) +
		"\n" + songsBaseStyle.Render(m.songs.View()) +
		"\n (↑/↓) navigate, (tab) switch between lists (q) quit\n"
}
