package tui

import (
	"fmt"
	"strings"
	bubbletea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) Init() bubbletea.Cmd {
	return nil
}

func (m Model) View() string {
    var s string

    switch m.step {

    case 0:
        // Header & search
        s += titleStyle.Render("📇 WhatsApp TUI\n\n")
        s += searchStyle.Render(fmt.Sprintf("Search: %s|", m.searchQuery)) + "\n\n"

        if len(m.filteredContacts) == 0 {
            s += normalStyle.Render("No matching contacts found.\n")
        } else {
            start := m.scrollOffset
            end := start + m.windowHeight
            if end > len(m.filteredContacts) {
                end = len(m.filteredContacts)
            }

            // Extract exactly the slice of contacts to show
            visible := m.filteredContacts[start:end]

            // Loop through visible contacts
            for i, c := range visible {
                absoluteIdx := start + i
                if absoluteIdx == m.selected {
                    // Highlight the selected contact
                    s += selectedStyle.Render(
                        fmt.Sprintf("-> %s (%s)", c.Name, c.Number),
                    ) + "\n"
                }
            }

            // If there are more contacts off‐screen, show a little pager hint
            if len(m.filteredContacts) > m.windowHeight {
                shown := end - start
                total := len(m.filteredContacts)
                s += normalStyle.Render(fmt.Sprintf("(%d/%d contacts)", shown, total)) + "\n"
            }
        }

        s += helpStyle.Render("↑/↓: Navigate | Enter: Select | Esc: Clear search | q: Quit")


		case 1:
				// Render the title
				s += titleStyle.Render(fmt.Sprintf("💬 Message to %s\n\n", m.filteredContacts[m.selected].Name))

				// Ensure messageLines has at least two lines
				if len(m.messageLines) < 2 {
						m.messageLines = []string{"", ""}
						m.currentLine = 1
				}

				// Render message lines with cursor
				var messageLinesRendered []string
				for i, line := range m.messageLines {
						if i == m.currentLine {
								// Render the current line with a styled cursor
								cursor := lipgloss.NewStyle().
										Foreground(lipgloss.Color("#00FFFF")).
										Render("▋")
								messageLinesRendered = append(messageLinesRendered, line+cursor)
						} else {
								messageLinesRendered = append(messageLinesRendered, line)
						}
				}

				// Join the lines and render with message style
				renderedMessage := strings.Join(messageLinesRendered, "\n")
				s += messageStyle.Render(fmt.Sprintf("│ %s\n", renderedMessage))

				// Render help instructions
				s += helpStyle.Render("Enter: Send | Ctrl+N: New line | Backspace: Delete | q: Quit")

    case 2:
        if strings.HasPrefix(m.messageLines[0], "✅") {
            s += successStyle.Render(strings.Join(m.messageLines, "\n"))
        } else {
            s += errorStyle.Render(strings.Join(m.messageLines, "\n"))
        }
        s += "\n\n"

        s += helpStyle.Render("Enter: Continue | q: Quit")
    }

    if m.isSending {
        s += "\n" + statusStyle.Render("⏳ " + m.statusMessage)
    } else if m.statusMessage != "" {
        s += "\n" + statusStyle.Render(m.statusMessage)
    }

    return s + "\n"
}


// sendMessageCmd is a command to send a message asynchronously.
type sendMessageResult struct {
	err error
}



