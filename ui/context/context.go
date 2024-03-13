// Package context contains program wide context and styling.
package context

import "github.com/ffalor/tasky/ui/theme"

// ProgramContext track information about the tui.
type ProgramContext struct {
	Theme             theme.Theme
	Styles            Styles
	ScreenHeight      int
	ScreenWidth       int
	MainContentHeight int
}
