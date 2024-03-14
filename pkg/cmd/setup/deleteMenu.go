package setup

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var deleteDocStyle = lipgloss.NewStyle()

type deleteActionChoice string

const (
	deleteYes deleteActionChoice = "yes"
	deleteNo  deleteActionChoice = "no"
)

type deleteItem struct {
	choice      deleteActionChoice
	label       string
	description string
}

func (i deleteItem) Title() string {
	return i.label
}

func (i deleteItem) Description() string {
	return i.description
}

func (i deleteItem) ID() string {
	return string(i.choice)
}

func (i deleteItem) FilterValue() string {
	return i.label
}

type deleteModel struct {
	list     list.Model
	selected *deleteActionChoice
}

func (m deleteModel) Init() tea.Cmd {
	return nil
}

func (m deleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			*m.selected = deleteNo
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := deleteDocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.list, cmd = m.list.Update(msg)
	*m.selected = m.list.SelectedItem().(deleteItem).choice

	return m, cmd
}

func (m deleteModel) View() string {
	return deleteDocStyle.Render(m.list.View())
}

func deleteMenu() error {
	items := []list.Item{
		deleteItem{choice: deleteYes, label: "Yes", description: "Delete the configuration"},
		deleteItem{choice: deleteNo, label: "No", description: "Do not delete the configuration"},
	}

	selected := deleteNo

	m := deleteModel{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		selected: &selected,
	}
	m.list.Title = "Are you sure you want to delete the configuration?"

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	if *m.selected == deleteYes {
		return os.Remove(".env")
	}

	return nil
}
