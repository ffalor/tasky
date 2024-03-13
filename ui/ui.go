// Package ui implements the termianl user interface.
package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/ffalor/tasky/ui/common"
	"github.com/ffalor/tasky/ui/context"
	"github.com/ffalor/tasky/ui/keys"
	"github.com/ffalor/tasky/ui/statusbar"
	"github.com/ffalor/tasky/ui/tabs"
	"github.com/ffalor/tasky/ui/theme"
)

// initMsg signals the start of the tui.
type initMsg struct{}

// Model the main ui model.
type Model struct {
	ctx       context.ProgramContext
	keys      keys.KeyMap
	statusBar statusbar.Model
	tabs      tabs.Model
}

// NewModel creates the main ui model with defaults.
func NewModel() Model {
	m := Model{
		keys: keys.Keys,
	}

	m.statusBar = statusbar.NewModel(m.ctx)
	m.tabs = tabs.NewModel()

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
				m.ctx.MainContentHeight = m.ctx.MainContentHeight + common.StatusBarHeight - common.ExpandedHelpHeight
			} else {
				m.ctx.MainContentHeight = m.ctx.MainContentHeight + common.ExpandedHelpHeight - common.StatusBarHeight
			}
		case key.Matches(msg, m.keys.NextTab):
			m.tabs.NextTab()
		case key.Matches(msg, m.keys.PrevTab):
			m.tabs.PrevTab()
		}

	case initMsg:
		m.ctx.Theme = theme.DefaultTheme
		m.ctx.Styles = context.BuildStyles(m.ctx.Theme)
	}

	m.updateProgramContext()
	m.statusBar, statusBarCmd = m.statusBar.Update(msg)

	cmds = append(cmds, cmd, statusBarCmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.tabs.View(m.ctx))
	s.WriteString("\n")
	mainContent := lipgloss.NewStyle().Height(m.ctx.MainContentHeight).Render("No Main Content")
	s.WriteString(mainContent)
	s.WriteString("\n")
	s.WriteString(m.statusBar.View())
	return s.String()
}

// onWindowSizeChange update size sensitive values.
func (m *Model) onWindowSizeChange(msg tea.WindowSizeMsg) {
	m.statusBar.SetWidth(msg.Width)
	m.ctx.ScreenHeight = msg.Height
	m.ctx.ScreenWidth = msg.Width

	if m.statusBar.ShowAll {
		m.ctx.MainContentHeight = msg.Height - common.TabsHeight - common.ExpandedHelpHeight
	} else {
		m.ctx.MainContentHeight = msg.Height - common.TabsHeight - common.StatusBarHeight
	}
}

// updateProgramContext updates the ProgramContext across our different models.
func (m *Model) updateProgramContext() {
	m.statusBar.UpdateProgramContext(&m.ctx)
}
