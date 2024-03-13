// Package context contains program wide context and styling.
package context

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
	"github.com/ffalor/tasky/ui/common"
	"github.com/ffalor/tasky/ui/theme"
)

// Styles all styles used by tasky.
type Styles struct {
	Colors struct {
		SuccessText lipgloss.AdaptiveColor
	}

	MainTextStyle  lipgloss.Style
	StatusBarStyle lipgloss.Style
	HelpStyle      help.Styles

	Tabs struct {
		Tab          lipgloss.Style
		ActiveTab    lipgloss.Style
		TabSeparator lipgloss.Style
		TabsRow      lipgloss.Style
	}
}

// BuildStyles builds the default styles to be used across the tui.
func BuildStyles(theme theme.Theme) Styles {
	var s Styles

	s.Colors.SuccessText = lipgloss.AdaptiveColor{
		Light: "#3DF294",
		Dark:  "#3DF294",
	}

	s.HelpStyle = help.Styles{
		ShortDesc:      lipgloss.NewStyle().Foreground(theme.FaintText),
		FullDesc:       lipgloss.NewStyle().Foreground(theme.FaintText),
		ShortSeparator: lipgloss.NewStyle().Foreground(theme.SecondaryBorder),
		FullSeparator:  lipgloss.NewStyle().Foreground(theme.SecondaryText),
		FullKey:        lipgloss.NewStyle().Foreground(theme.PrimaryText),
		ShortKey:       lipgloss.NewStyle().Foreground(theme.PrimaryText),
		Ellipsis:       lipgloss.NewStyle().Foreground(theme.SecondaryText),
	}

	s.MainTextStyle = lipgloss.NewStyle().
		Foreground(theme.PrimaryText).
		Bold(true)

	s.StatusBarStyle = lipgloss.NewStyle().
		Background(theme.SelectedBackground).
		Height(common.StatusBarHeight)

	s.Tabs.Tab = lipgloss.NewStyle().
		Faint(true).
		Padding(0, 2)

	s.Tabs.ActiveTab = s.Tabs.Tab.Copy().
		Faint(false).
		Bold(true).
		Background(theme.SelectedBackground).
		Foreground(theme.PrimaryText)

	s.Tabs.TabSeparator = lipgloss.NewStyle().
		Foreground(theme.SecondaryBorder)

	s.Tabs.TabsRow = lipgloss.NewStyle().
		Height(common.TabsContentHeight).
		PaddingTop(1).
		PaddingBottom(0).
		BorderBottom(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderBottomForeground(theme.PrimaryBorder)

	return s
}
