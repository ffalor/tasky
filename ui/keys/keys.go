// Package keys impelements keys used by the tui.
package keys

import "github.com/charmbracelet/bubbles/key"

// KeyMap all program keys.
type KeyMap struct {
	AddTask key.Binding
	Start   key.Binding
	Reset   key.Binding
	Quit    key.Binding
	Help    key.Binding
}

// FullHelp returns the FullHelp for all program keys.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.AddTask,
			k.Start,
			k.Reset,
			k.Quit,
			k.Help,
		},
	}
}

// ShortHelp returns the shorthelp which contains only the help key.
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Help,
	}
}

// Keys program keys with their help bindings.
var Keys = KeyMap{
	AddTask: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "Add a new task"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Start: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "start"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "reset"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
}
