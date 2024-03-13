// Package keys impelements keys used by the tui.
package keys

import "github.com/charmbracelet/bubbles/key"

// KeyMap all program keys.
type KeyMap struct {
	AddTask key.Binding
	Quit    key.Binding
	NextTab key.Binding
	PrevTab key.Binding
	Help    key.Binding
}

// FullHelp returns the FullHelp for all program keys.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			k.AddTask,
			k.NextTab,
			k.PrevTab,
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
	NextTab: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "next tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "previous tab"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
}
