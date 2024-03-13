package setup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sirupsen/logrus"
)

type actionChoice int

const (
	actionNone actionChoice = iota
	actionCreate
	actionUpdate
	actionDelete
	actionView
)

type actionModel struct {
	actions  []actionChoice
	cursor   int
	selected actionChoice
}

func newActionModel() *actionModel {
	return &actionModel{
		actions: []actionChoice{
			actionCreate,
			actionUpdate,
			actionDelete,
			actionView,
		},
		selected: actionNone,
	}
}

func (m *actionModel) Init() tea.Cmd {
	return nil
}

func (m *actionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.actions)-1 {
				m.cursor++
			}
		case "enter", " ":
			if m.cursor < 0 || m.cursor >= len(m.actions) {
				m.cursor = 0
			} else {
				m.selected = m.actions[m.cursor]

				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m *actionModel) View() string {
	s := "\n"
	for i, choice := range m.actions {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		var out string
		switch choice {
		case actionCreate:
			out = "Create a new config"
		case actionUpdate:
			out = "Update existing config"
		case actionDelete:
			out = "Delete existing config"
		case actionView:
			out = "View existing config"
		}

		s += fmt.Sprintf("%s %s\n", cursor, out)
	}

	return s
}

func actionMenu() actionChoice {
	aModel := newActionModel()
	_, err := tea.NewProgram(aModel).Run()
	if err != nil {
		logrus.WithError(err).Fatal("encountered an error while running the TUI")
	}

	return aModel.selected
}
