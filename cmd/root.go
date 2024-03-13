// Package cmd is the tui entrypoint. CLI parameters and commonds that modify the startup of the tui are added here.
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ffalor/tasky/ui"
	"github.com/muesli/termenv"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tasky",
		Short: "A simple tui task tracker",
	}
)

// Execute main entry point into tasky tui
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func createModel() ui.Model {
	return ui.NewModel()
}

func init() {
	rootCmd.Run = func(_ *cobra.Command, _ []string) {
		lipgloss.SetHasDarkBackground(termenv.HasDarkBackground())

		model := createModel()
		p := tea.NewProgram(
			model,
			tea.WithAltScreen(),
		)

		if _, err := p.Run(); err != nil {
			fmt.Printf("Failed to start the TUI: %s", err)
			os.Exit(1)
		}
	}
}
