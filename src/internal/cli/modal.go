package cli

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().BorderStyle(lipgloss.InnerHalfBlockBorder()).BorderForeground(lipgloss.Color("240"))

type Model struct {
	table table.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s price %s!", m.table.SelectedRow()[0], m.table.SelectedRow()[2]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func New(rows [][]string) Model {
	tableColumns := []table.Column{
		{Title: "Name", Width: 15},
		{Title: "Location", Width: 10},
		{Title: "Price", Width: 10},
		{Title: "Status", Width: 10},
	}
	var tableRows []table.Row
	for _, record := range rows {
		dict := []string{
			record[0],
			record[1],
			record[2],
			record[3],
		}
		tableRows = append(tableRows, dict)
	}

	newTable := table.New(
		table.WithColumns(tableColumns),
		table.WithRows(tableRows),
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
	newTable.SetStyles(s)

	m := Model{table: newTable}
	return m
}
