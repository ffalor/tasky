// Package ui implements the termianl user interface.
package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ffalor/tasky/ui/common"
	"github.com/ffalor/tasky/ui/context"
	"github.com/ffalor/tasky/ui/keys"
	"github.com/ffalor/tasky/ui/statusbar"
	"github.com/ffalor/tasky/ui/table"
	"github.com/ffalor/tasky/ui/theme"
)

// initMsg signals the start of the tui.
type initMsg struct{}

// Model the main ui model.
type Model struct {
	ctx       context.ProgramContext
	statusBar statusbar.Model
	table     table.Model
	keys      keys.KeyMap
}

// NewModel creates the main ui model with defaults.
func NewModel() Model {
	m := Model{
		keys: keys.Keys,
	}

	statusBar := statusbar.NewModel(m.ctx)
	m.statusBar = statusBar

	table := table.New(m.ctx)
	m.table = table

	return m
}

func (m *Model) initTUI() tea.Msg {
	return initMsg{}
}

func (m Model) Init() tea.Cmd {
	return m.initTUI
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd          tea.Cmd
		statusBarCmd tea.Cmd
		tableCmd     tea.Cmd
		cmds         []tea.Cmd
	)

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.onWindowSizeChange(msg)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			// todo replace use main content to track size
			if !m.statusBar.ShowAll {
				m.ctx.ScreenHeight = m.ctx.ScreenHeight + common.StatusBarHeight - common.ExpandedHelpHeight
			} else {
				m.ctx.ScreenHeight = m.ctx.ScreenHeight + common.ExpandedHelpHeight - common.StatusBarHeight
			}
		}

	case initMsg:
		m.ctx.Theme = theme.DefaultTheme
		m.ctx.Styles = context.BuildStyles(m.ctx.Theme)
	}

	m.updateProgramContext()
	m.statusBar, statusBarCmd = m.statusBar.Update(msg)
	m.table, tableCmd = m.table.Update(msg)

	cmds = append(cmds, cmd, statusBarCmd, tableCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	t := lipgloss.JoinVertical(lipgloss.Top, m.table.View(), m.statusBar.View())
	return t
}

// onWindowSizeChange update size sensitive values.
func (m *Model) onWindowSizeChange(msg tea.WindowSizeMsg) {
	m.table.SetTableHeight(msg.Height)
	m.statusBar.SetWidth(msg.Width)
	m.ctx.ScreenHeight = msg.Height
	m.ctx.ScreenWidth = msg.Width
}

// updateProgramContext updates the ProgramContext across our different models.
func (m *Model) updateProgramContext() {
	m.statusBar.UpdateProgramContext(&m.ctx)
	m.table.UpdateProgramContext(&m.ctx)
}
