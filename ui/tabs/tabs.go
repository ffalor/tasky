package tabs

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ffalor/tasky/ui/context"
)

type Model struct {
	tabTitles        []string
	CurrentSectionId int
}

// NewModel create a new model with defaults for the tabs ui component
func NewModel() Model {
	return Model{
		tabTitles:        []string{"Open", "Closed", "Blocked"},
		CurrentSectionId: 0,
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View(ctx context.ProgramContext) string {

	var tabs []string
	for i, tabTitle := range m.tabTitles {
		if m.CurrentSectionId == i {
			tabs = append(tabs, ctx.Styles.Tabs.ActiveTab.Render(tabTitle))
		} else {
			tabs = append(tabs, ctx.Styles.Tabs.Tab.Render(tabTitle))
		}
	}

	renderedTabs := lipgloss.NewStyle().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, strings.Join(tabs, ctx.Styles.Tabs.TabSeparator.Render("|"))))

	return ctx.Styles.Tabs.TabsRow.Copy().
		Width(ctx.ScreenWidth).
		MaxWidth(ctx.ScreenWidth).
		Render(renderedTabs)
}

func (m *Model) NextTab() {
	m.CurrentSectionId = (m.CurrentSectionId + 1) % len(m.tabTitles)
}

func (m *Model) PrevTab() {
	m.CurrentSectionId = (m.CurrentSectionId - 1 + len(m.tabTitles)) % len(m.tabTitles)
}
