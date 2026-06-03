package components

import (
	"github.com/charmbracelet/bubbletea"
)

type CursorList struct {
	Items      []string
	Cursor     int
	Reverse    bool
	OnPressMap map[string]func()
}

func (c *CursorList) Init() tea.Cmd {
	return nil
}

func (c *CursorList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		default:
			if action, ok := c.OnPressMap[key]; ok {
				action()
				break
			}

			switch key {
			case "down":
				if c.Reverse {
					c.Dec()
				} else {
					c.Inc()
				}
			case "up":
				if c.Reverse {
					c.Inc()
				} else {
					c.Dec()
				}
			}
		}
	}

	return c, nil
}

func (c *CursorList) View() string {
	return ""
}

func (c *CursorList) Inc() {
	if c.Cursor < len(c.Items)-1 {
		c.Cursor++
	}
}

func (c *CursorList) Dec() {
	if c.Cursor > 0 {
		c.Cursor--
	}
}

func (c *CursorList) GetCursor() int {
	return c.Cursor
}
