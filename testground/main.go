package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table    table.Model
	selected string // Store the ID of the selected account
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			// THIS IS YOUR "MENU" SELECTION LOGIC
			m.selected = m.table.SelectedRow()[1] // Get ID from the second column
			return m, nil
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	s := baseStyle.Render(m.table.View()) + "\n"
	if m.selected != "" {
		s += fmt.Sprintf("\n  Selected Account ID: %s\n", m.selected)
	} else {
		s += "\n  Press Enter to select an account."
	}
	return s
}

func main() {
	columns := []table.Column{
		{Title: "Account Name", Width: 20},
		{Title: "ID", Width: 12},
		{Title: "Status", Width: 10},
	}

	rows := []table.Row{
		{"Production", "123456789", "Active"},
		{"Staging", "987654321", "Active"},
		{"Development", "112233445", "Suspended"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true), // Essential for keyboard nav
		table.WithHeight(7),
	)

	// Custom Styling
	s := table.DefaultStyles()
	s.Header = s.Header.BorderStyle(lipgloss.NormalBorder()).Bold(true)
	s.Selected = s.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57")).Bold(true)
	t.SetStyles(s)

	m := model{table: t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
