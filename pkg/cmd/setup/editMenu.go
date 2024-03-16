package setup

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
)

type item struct {
	label string
	value string
}

type editModel struct {
	itemCursor int
	editCursor int
	items      []item
	editting   bool
}

var cursorFmt = color.New(color.FgBlack, color.BgWhite).SprintFunc()

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
			if !m.editting {
				m.itemCursor = max(m.itemCursor-1, 0)
			}
		case "down":
			if !m.editting {
				m.itemCursor = min(m.itemCursor+1, len(m.items)-1)
			}
		case "left":
			if m.editting {
				m.editCursor = max(m.editCursor-1, 0)
			}
		case "right":
			if m.editting {
				m.editCursor = min(m.editCursor+1, len(m.items[m.itemCursor].value))
			}
		case "enter":
			m.editting = !m.editting
			m.editCursor = len(m.items[m.itemCursor].value)
		case "ctrl+c":
			return m, tea.Quit
		case "backspace", "delete":
			if m.editting {
				if m.editCursor < len(m.items[m.itemCursor].value) {
					m.items[m.itemCursor].value = m.items[m.itemCursor].value[:m.editCursor-1] + m.items[m.itemCursor].value[m.editCursor:]
				} else if m.editCursor == len(m.items[m.itemCursor].value) {
					m.items[m.itemCursor].value = m.items[m.itemCursor].value[:m.editCursor-1]
					m.editCursor--
				}
			}
		case "esc":
			if m.editting {
				m.editting = false
			} else {
				return m, tea.Quit
			}
		default:
			if m.editting {
				m.items[m.itemCursor].value += string(msg.Runes)
				m.editCursor++
			} else {
				switch msg.String() {
				case "j":
					m.itemCursor = min(m.itemCursor+1, len(m.items)-1)
				case "k":
					m.itemCursor = max(m.itemCursor-1, 0)
				case "q":
					return m, tea.Quit
				}
			}
		}
	}
	return m, nil
}

func (m editModel) View() string {
	s := ""
	for i, item := range m.items {
		prefix := "[ ]"
		if i == m.itemCursor {
			if m.editting {
				prefix = "[*]"
			} else {
				prefix = "[>]"
			}
		}

		if m.editting && i == m.itemCursor {
			if m.editCursor == len(item.value) {
				s += fmt.Sprintf("%s %s: %s%s\n", prefix, item.label, item.value, cursorFmt(" "))
			} else if m.editCursor == len(item.value)-1 {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value[:m.editCursor]+cursorFmt(string(item.value[m.editCursor])))
			} else {
				s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value[:m.editCursor]+cursorFmt(string(item.value[m.editCursor]))+item.value[m.editCursor+1:])
			}

		} else {
			s += fmt.Sprintf("%s %s: %s\n", prefix, item.label, item.value)
		}
	}
	return s
}

func editMenu() error {
	p := tea.NewProgram(newEditModel())
	_, err := p.Run()
	return err
}
