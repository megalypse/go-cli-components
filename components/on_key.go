package components

import (
	"slices"

	"github.com/charmbracelet/bubbletea"
)

type OnKeyFn func(key string) (tea.Model, tea.Cmd)

func OnKey(action func() (tea.Model, tea.Cmd), keys ...string) func(key string) (tea.Model, tea.Cmd) {
	return func(key string) (tea.Model, tea.Cmd) {
		if slices.Contains(keys, key) {
			return action()
		}

		return nil, nil
	}
}
