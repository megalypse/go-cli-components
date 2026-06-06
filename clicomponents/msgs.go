package clicomponents

import tea "github.com/charmbracelet/bubbletea"

func MsgPublish[T any](msg T) func() tea.Msg {
	return func() tea.Msg {
		return msg
	}
}

type MsgEditModeUpdate struct {
	Mode bool
}
