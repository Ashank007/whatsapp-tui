package tui

import (
	"fmt"
	"strings"
	"whatsapp-tui/contacts"	
	bubbletea "github.com/charmbracelet/bubbletea"
	waTypes "go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"context"
)

func (m *Model) adjustScroll() {
	if m.selected < m.scrollOffset {
		m.scrollOffset = m.selected
	}
	if m.selected >= m.scrollOffset+m.windowHeight {
		m.scrollOffset = m.selected - m.windowHeight + 1
	}
}

func (m *Model) adjustSelection() {
	if m.selected >= len(m.filteredContacts) && len(m.filteredContacts) > 0 {
		m.selected = len(m.filteredContacts) - 1
	} else if len(m.filteredContacts) == 0 {
		m.selected = 0
	}
	m.scrollOffset = 0
	m.adjustScroll()
}

func (m *Model) filterContacts() {
	if m.searchQuery == "" {
		m.filteredContacts = m.contacts
		return
	}
	query := strings.ToLower(m.searchQuery)
	filtered := []contacts.Contact{}
	for _, c := range m.contacts {
		if strings.Contains(strings.ToLower(c.Name), query) || strings.Contains(strings.ToLower(c.Number), query) {
			filtered = append(filtered, c)
		}
	}
	m.filteredContacts = filtered
}

func (m Model) sendMessageCmd() bubbletea.Cmd {
    return func() bubbletea.Msg {
        if len(m.filteredContacts) == 0 || m.selected >= len(m.filteredContacts) {
            return sendMessageResult{err: fmt.Errorf("no contact selected")}
        }
        number := m.filteredContacts[m.selected].Number
        recipient := waTypes.NewJID(number, "s.whatsapp.net")
        // Only include content from second line onward
        fullMessage := strings.Join(m.messageLines[1:], "\n")
        msg := &waE2E.Message{
            Conversation: &fullMessage,
        }
        _, err := m.client.SendMessage(context.Background(), recipient, msg)
        return sendMessageResult{err: err}
    }
}
