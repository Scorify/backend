package setup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type item struct {
	label   string
	value   string
	editing bool
}

type editModel struct {
	cursor int
	items  []item
}

func newEditModel() editModel {
	return editModel{
		items: []item{
			{label: "Name", value: "John Doe"},
			{label: "Email", value: "john.doe@example.com"},
		},
	}
}

func (m editModel) Init() tea.Cmd {
	return nil
}

func (m editModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			m.items[m.cursor].editing = !m.items[m.cursor].editing
		case "ctrl+c":
			return m, tea.Quit
		default:
			if m.items[m.cursor].editing {
				m.items[m.cursor].value += string(msg.Runes)
			}
		}
	}
	return m, nil
}

func (m editModel) View() string {
	s := ""
	for i, item := range m.items {
		prefix := "  "
		if i == m.cursor {
			prefix = "> "
		}
		if item.editing {
			s += fmt.Sprintf("%s*%s*: %s (editing)\n", prefix, item.label, item.value)
		} else {
			s += fmt.Sprintf("%s%s: %s\n", prefix, item.label, item.value)
		}
	}
	return s
}

func editMenu() error {
	p := tea.NewProgram(newEditModel())
	_, err := p.Run()
	return err
}
