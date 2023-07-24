package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type sign int

const (
	X sign = iota
	O
)

const size = 3

type selected = map[coordinate]sign

type coordinate struct {
	x int
	y int
}

type model struct {
	cursor   coordinate
	board    [][]string
	selected selected
}

func newModel() model {
	return model{
		board: [][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
		selected: make(selected),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	less := func(c, max int) int {
		if c > 0 {
			return c - 1
		} else {
			return max
		}
	}

	more := func(c, min int) int {
		if c < size-1 {
			return c + 1
		} else {
			return min
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			m.cursor.y = less(m.cursor.y, size-1)
		case "down", "j":
			m.cursor.y = more(m.cursor.y, 0)
		case "left", "h":
			m.cursor.x = less(m.cursor.x, size-1)
		case "right", "l":
			m.cursor.x = more(m.cursor.x, 0)
		case "enter", " ":
			m.selected[m.cursor] = X
		}
	}

	return m, nil
}

func (m model) getSelected(j, i int) (string, error) {
	if s, ok := m.selected[coordinate{x: j, y: i}]; ok {
		switch s {
		case X:
			return "X", nil
		case O:
			return "O", nil
		}
	}

	return ".", nil
}

func (m model) View() string {
	ret := "Let's play a game!\n"

	for i, row := range m.board {
		for j, col := range row {
			if s, err := m.getSelected(j, i); err == nil {
				col = s
			}
			if m.cursor.x == j && m.cursor.y == i {
				ret += fmt.Sprintf("[%s]", col)
			} else {
				ret += fmt.Sprintf(" %s ", col)
			}
		}
		ret += "\n"
	}

	ret += "\nPress q or ctrl+c to quit\n"
	return ret
}

func main() {
	fmt.Println("Welcome to the game of Tic Tac Toe!")

	p := tea.NewProgram(newModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error while runninng the program; %v", err)
		os.Exit(1)
	}
}
