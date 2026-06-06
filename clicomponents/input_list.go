package clicomponents

import tea "github.com/charmbracelet/bubbletea"

type Focusable interface {
	Focus() tea.Cmd
	Blur()
}

type Updatable interface {
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
}

type InputItem interface {
	Focusable
	Updatable
}

type InputList struct {
	CurrentInput   int
	Inputs         []InputItem
	EditMode       bool
	EditModeOnKeys []OnKeyFn
	CmdModeOnKeys  []OnKeyFn
}

func (i *InputList) Init() tea.Cmd {
	return nil
}

func (i *InputList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case MsgEditModeUpdate:
		i.EditMode = msg.Mode
		if i.EditMode {
			cmds = append(cmds, i.Inputs[i.CurrentInput].Focus())
		}
	}

	_, cmd := func() (tea.Model, tea.Cmd) {
		if i.EditMode {
			return i.handleEditModeCmd(msg)
		}

		return i.handleCmdModeCmd(msg)

	}()

	cmds = append(cmds, cmd)
	for _, item := range i.Inputs {
		_, cmd := item.Update(msg)
		cmds = append(cmds, cmd)
	}

	return i, tea.Batch(cmds...)
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

		for _, fn := range []OnKeyFn{
			OnKey("tab", func() (tea.Model, tea.Cmd) {
				if i.CurrentInput == len(i.Inputs)-1 {
					i.CurrentInput = 0
				} else {
					i.CurrentInput++
				}

				i.unfocusAll()
				return nil, i.Inputs[i.CurrentInput].Focus()
			}),

			OnKey("shift+tab", func() (tea.Model, tea.Cmd) {
				if i.CurrentInput == 0 {
					i.CurrentInput = len(i.Inputs) - 1
				} else {
					i.CurrentInput--
				}

				i.unfocusAll()
				return nil, i.Inputs[i.CurrentInput].Focus()
			}),

			OnKey("esc", func() (tea.Model, tea.Cmd) {
				i.unfocusAll()
				return nil, MsgPublish(MsgEditModeUpdate{false})
			}),
		} {
			_, cmd := fn(key)
			if cmd != nil {
				return nil, cmd
			}
		}

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

		if _, cmd := OnKey("i", func() (tea.Model, tea.Cmd) {
			return nil, MsgPublish(MsgEditModeUpdate{true})
		})(key); cmd != nil {
			return nil, cmd
		}
	}

	return nil, nil
}

func (i *InputList) unfocusAll() {
	for _, input := range i.Inputs {
		input.Blur()
	}
}
