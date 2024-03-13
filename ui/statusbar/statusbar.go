// Package statusbar impelments the statusbar that appears at the bottom of the tui.
package statusbar

import (
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ffalor/tasky/ui/context"
	"github.com/ffalor/tasky/ui/keys"
	"github.com/ffalor/tasky/ui/theme"
)

// temp will be removed
var StatusStyle = lipgloss.NewStyle().
	Foreground(theme.DefaultTheme.InvertedText).
	Background(theme.DefaultTheme.SecondaryText).
	Padding(0, 1).
	Bold(true)

// Model the statusbar model.
type Model struct {
	ctx          *context.ProgramContext
	leftSection  *string
	rightSection *string
	help         help.Model
	ShowAll      bool
}

// NewModel creates the statusbar model with defaults.
func NewModel(ctx context.ProgramContext) Model {
	help := help.New()
	help.ShowAll = true
	l := ""
	r := ""
	return Model{
		ctx:          &ctx,
		help:         help,
		leftSection:  &l,
		rightSection: &r,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.Help):
			m.ShowAll = !m.ShowAll
		}
	}

	return m, nil
}

func (m Model) View() string {
	helpIndicator := lipgloss.NewStyle().
		Background(theme.DefaultTheme.FaintText).
		Foreground(theme.DefaultTheme.SelectedBackground).
		Padding(0, 1).
		Render("? help")
	viewSwitcher := StatusStyle.Render("🗊 Tasks")
	leftSection := ""
	if m.leftSection != nil {
		leftSection = *m.leftSection
	}
	rightSection := ""
	if m.rightSection != nil {
		rightSection = *m.rightSection
	}
	spacing := lipgloss.NewStyle().
		Background(theme.DefaultTheme.SelectedBackground).
		Render(
			strings.Repeat(
				" ",
				max(0,
					m.ctx.ScreenWidth-lipgloss.Width(
						viewSwitcher,
					)-lipgloss.Width(leftSection)-
						lipgloss.Width(rightSection)-
						lipgloss.Width(
							helpIndicator,
						),
				)))

	statusbar := m.ctx.Styles.StatusBarStyle.Copy().
		Render(lipgloss.JoinHorizontal(lipgloss.Top, viewSwitcher, leftSection, spacing, rightSection, helpIndicator))

	if m.ShowAll {
		fullHelp := m.help.View(keys.Keys)
		return lipgloss.JoinVertical(lipgloss.Top, statusbar, fullHelp)
	}

	return statusbar
}

// SetWidth sets the width of the statusbar.
func (m *Model) SetWidth(width int) {
	m.help.Width = width
}

func (m *Model) UpdateProgramContext(ctx *context.ProgramContext) {
	m.ctx = ctx
	m.help.Styles = m.ctx.Styles.HelpStyle
}

// SetLeftSection sets the left section of the statusbar.
func (m *Model) SetLeftSection(leftSection string) {
	*m.leftSection = leftSection
}

// SetRightSection sets the right section of the statusbar.
func (m *Model) SetRightSection(rightSection string) {
	*m.rightSection = rightSection
}
