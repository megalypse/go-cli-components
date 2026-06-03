package components

import (
	"strings"

	"github.com/charmbracelet/bubbletea"
)

type CursorListVertical struct {
	Items      []string
	Cursor     int
	KeyActions []OnKeyFn
	RenderSize int
}

func (c *CursorListVertical) Init() tea.Cmd {
	return nil
}

func (c *CursorListVertical) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		for _, keyAction := range c.KeyActions {
			model, cmd := keyAction(msg.String())
			if model != nil || cmd != nil {
				return model, cmd
			}
		}

		switch msg.String() {
		case "up":
			c.Dec()
		case "down":
			c.Inc()
		}
	}

	return nil, nil
}

func (c *CursorListVertical) View() string {
	startDelta := c.RenderSize - c.Cursor
	endDelta := (c.Cursor + 1) + c.RenderSize - len(c.Items)

	listBuilder := strings.Builder{}

	for range startDelta {
		listBuilder.WriteString("\n")
	}

	for i, item := range c.Items {
		isPreviousOffLimits := c.RenderSize > 0 && i < c.Cursor-c.RenderSize
		isNextOffLimits := c.RenderSize > 0 && i > c.Cursor+c.RenderSize
		if isPreviousOffLimits || isNextOffLimits {

			continue
		}

		if i == c.Cursor {
			listBuilder.WriteString("> ")
		}
		listBuilder.WriteString(item)
		listBuilder.WriteString("\n")
	}

	for range endDelta {
		listBuilder.WriteString("\n")
	}

	return listBuilder.String()
}

func (c *CursorListVertical) Inc() {
	if c.Cursor < len(c.Items)-1 {
		c.Cursor++
	}
}

func (c *CursorListVertical) Dec() {
	if c.Cursor > 0 {
		c.Cursor--
	}
}

func (c *CursorListVertical) GetCursor() int {
	return c.Cursor
}
