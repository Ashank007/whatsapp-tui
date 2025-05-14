package tui

import ("github.com/charmbracelet/lipgloss")
// Styles for the TUI
var (
	// Define styles using lipgloss
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00")).
			Align(lipgloss.Left).
			Padding(0, 0).
			Width(80)

	selectedStyle = lipgloss.NewStyle().
	    Background(lipgloss.Color("#444444")).
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			Align(lipgloss.Left)

	normalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Align(lipgloss.Left)

	searchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFFF")).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#888888")).
			Align(lipgloss.Left).
			Padding(0, 1).
			Width(80)

	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")).
			Align(lipgloss.Left).
			Width(80)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true).
			Align(lipgloss.Left).
			Width(80)

	successStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true).
			Align(lipgloss.Left).
			Width(80)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			MarginTop(1).
			Align(lipgloss.Left).
			Width(80)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			MarginTop(1).
			Align(lipgloss.Left).
			Width(80)
)

