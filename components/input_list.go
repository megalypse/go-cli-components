package components

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	Focus() tea.Cmd
	Blur()
}

type InputList struct {
	CurrentInput   int
	Inputs         []Focusable
	EditMode       bool
	EditModeOnKeys []OnKeyFn
	CmdModeOnKeys  []OnKeyFn
}

func (i *InputList) Init() tea.Cmd {

	return nil
}

func (i *InputList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if i.EditMode {
		return i.handleEditModeCmd(msg)
	}

	return i.handleCmdModeCmd(msg)
}

func (i *InputList) View() string {
	return ""
}

func (i *InputList) handleEditModeCmd(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		for _, keyAction := range i.EditModeOnKeys {
			model, cmd := keyAction(key)
			if model != nil || cmd != nil {
				return model, cmd
			}
		}

		OnKey("tab", func() (tea.Model, tea.Cmd) {
			if i.CurrentInput == len(i.Inputs)-1 {
				i.CurrentInput = 0
			} else {
				i.CurrentInput++
			}

			return nil, i.Inputs[i.CurrentInput].Focus()
		})(key)

		OnKey("shift+tab", func() (tea.Model, tea.Cmd) {
			if i.CurrentInput == 0 {
				i.CurrentInput = len(i.Inputs) - 1
			} else {
				i.CurrentInput--
			}

			return nil, i.Inputs[i.CurrentInput].Focus()
		})

		OnKey("esc", func() (tea.Model, tea.Cmd) {
			i.EditMode = false
			for _, i := range i.Inputs {
				i.Blur()
			}

			return nil, nil
		})

	}

	return nil, nil
}

func (i *InputList) handleCmdModeCmd(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		for _, keyAction := range i.CmdModeOnKeys {
			model, cmd := keyAction(key)
			if model != nil || cmd != nil {
				return model, cmd
			}
		}

		OnKey("i", func() (tea.Model, tea.Cmd) {
			i.EditMode = true
			return nil, i.Inputs[i.CurrentInput].Focus()
		})(key)
	}
	return nil, nil
}
