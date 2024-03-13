package setup

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var actionDocStyle = lipgloss.NewStyle()

type actionChoice string

const (
	actionNone   actionChoice = ""
	actionCreate actionChoice = "create"
	actionUpdate actionChoice = "update"
	actionDelete actionChoice = "delete"
	actionView   actionChoice = "view"
)

type actionItem struct {
	choice      actionChoice
	label       string
	description string
}

func (i actionItem) Title() string {
	return i.label
}

func (i actionItem) Description() string {
	return i.description
}

func (i actionItem) ID() string {
	return string(i.choice)
}

func (i actionItem) FilterValue() string {
	return i.label + " " + i.description
}

type actionModel struct {
	list     list.Model
	selected *actionChoice
}

func (m actionModel) Init() tea.Cmd {
	return nil
}

func (m actionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			*m.selected = actionNone
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := actionDocStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.list, cmd = m.list.Update(msg)
	*m.selected = m.list.SelectedItem().(actionItem).choice

	return m, cmd
}

func (m actionModel) View() string {
	return actionDocStyle.Render(m.list.View())
}

func actionMenu() (actionChoice, error) {
	items := []list.Item{
		actionItem{choice: actionCreate, label: "Create", description: "Create a new configuration"},
		actionItem{choice: actionUpdate, label: "Update", description: "Update existing configuration"},
		actionItem{choice: actionDelete, label: "Delete", description: "Delete existing configuration"},
		actionItem{choice: actionView, label: "View", description: "View existing configuration"},
	}

	selected := actionNone

	m := actionModel{
		list:     list.New(items, list.NewDefaultDelegate(), 0, 0),
		selected: &selected,
	}
	m.list.Title = "Select action to setup Scorify Configuration"

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return actionNone, err
	}

	return *m.selected, nil
}
