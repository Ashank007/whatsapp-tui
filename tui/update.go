package tui

import (
	bubbletea "github.com/charmbracelet/bubbletea"
	"strings"
)

func (m *Model) Update(msg bubbletea.Msg) (bubbletea.Model, bubbletea.Cmd) {
	switch msg := msg.(type) {
	case bubbletea.WindowSizeMsg:
		m.windowHeight = msg.Height - 10
		m.terminalWidth = msg.Width
		titleStyle = titleStyle.Width(m.terminalWidth)
		selectedStyle = selectedStyle.Width(m.terminalWidth)
		normalStyle = normalStyle.Width(m.terminalWidth)
		searchStyle = searchStyle.Width(m.terminalWidth)
		messageStyle = messageStyle.Width(m.terminalWidth)
		errorStyle = errorStyle.Width(m.terminalWidth)
		successStyle = successStyle.Width(m.terminalWidth)
		helpStyle = helpStyle.Width(m.terminalWidth)
		statusStyle = statusStyle.Width(m.terminalWidth)
		return m, nil

	case bubbletea.KeyMsg:
		switch msg.Type {
		case bubbletea.KeyCtrlN:
			if m.step == 1 {
				m.messageLines = append(m.messageLines, "")
				m.currentLine++
				return m, nil
			}

		case bubbletea.KeyEnter:
			if m.step == 0 {
				if len(m.filteredContacts) == 0 {
					m.statusMessage = "No contacts to select"
					return m, nil
				}
				m.step = 1
				m.statusMessage = ""
				// Simulate Ctrl+N: Add a new line and move cursor to it
				m.messageLines = []string{"", ""} // Start with two lines
				m.currentLine = 1                 // Cursor on the second line
			} else if m.step == 1 {
				fullMessage := strings.TrimSpace(strings.Join(m.messageLines, "\n"))
				if fullMessage == "" {
					m.statusMessage = "Message cannot be empty"
					return m, nil
				}
				m.isSending = true
				m.statusMessage = "Sending message..."
				return m, m.sendMessageCmd()
			} else if m.step == 2 {
				m.step = 0
				m.messageLines = []string{""}
				m.currentLine = 0
				m.searchQuery = ""
				m.filteredContacts = m.contacts
				m.selected = 0
				m.scrollOffset = 0
				m.statusMessage = "Connected"
			}

		case bubbletea.KeyUp:
			if m.step == 0 && m.selected > 0 {
				m.selected--
				m.adjustScroll()
			}

		case bubbletea.KeyDown:
			if m.step == 0 && m.selected < len(m.filteredContacts)-1 {
				m.selected++
				m.adjustScroll()
			}

		case bubbletea.KeyEsc:
			if m.step == 0 {
				m.searchQuery = ""
				m.filteredContacts = m.contacts
				m.selected = 0
				m.scrollOffset = 0
			}

		case bubbletea.KeyBackspace:
			if m.step == 0 && len(m.searchQuery) > 0 {
				m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
				m.filterContacts()
				m.adjustSelection()
			} else if m.step == 1 {
				if len(m.messageLines) == 1 && len(m.messageLines[0]) == 0 {
					return m, nil
				}
				if len(m.messageLines[m.currentLine]) > 0 {
					m.messageLines[m.currentLine] = m.messageLines[m.currentLine][:len(m.messageLines[m.currentLine])-1]
				} else if m.currentLine > 0 {
					m.messageLines = m.messageLines[:m.currentLine]
					m.currentLine--
				}
				if len(m.messageLines) == 0 {
					m.messageLines = []string{""}
					m.currentLine = 0
				}
			}

		case bubbletea.KeySpace:
			if m.step == 0 {
				m.searchQuery += " "
				m.filterContacts()
				m.adjustSelection()
			} else if m.step == 1 {
				m.messageLines[m.currentLine] += " "
			}

		case bubbletea.KeyRunes:
			if m.step == 0 {
				m.searchQuery += string(msg.Runes)
				m.filterContacts()
				m.adjustSelection()
			} else if m.step == 1 {
				if len(m.messageLines) == 0 {
					m.messageLines = []string{""}
					m.currentLine = 0
				}
				m.messageLines[m.currentLine] += string(msg.Runes)
			}
		}

		if msg.String() == "q" {
			return m, bubbletea.Quit
		}

	case sendMessageResult:
		m.isSending = false
		if msg.err != nil {
			m.messageLines = []string{"❌ Failed to send: " + msg.err.Error()}
			m.statusMessage = "Error sending message"
		} else {
			m.messageLines = []string{"✅ Message sent!"}
			m.statusMessage = "Message sent successfully"
		}
		m.currentLine = 0
		m.step = 2
	}

	return m, nil
}
