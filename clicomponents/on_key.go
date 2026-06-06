package clicomponents

import (
	"slices"

	"github.com/charmbracelet/bubbletea"
)

type OnKeyFn func(key string) (tea.Model, tea.Cmd)

func OnKeys(action func() (tea.Model, tea.Cmd), keys ...string) func(key string) (tea.Model, tea.Cmd) {
	return func(key string) (tea.Model, tea.Cmd) {
		if slices.Contains(keys, key) {
			return action()
		}

		return nil, nil
	}
}

func OnKey(key string, action func() (tea.Model, tea.Cmd)) func(key string) (tea.Model, tea.Cmd) {
	return func(k string) (tea.Model, tea.Cmd) {
		if k == key {
			return action()
		}

		return nil, nil
	}
}
