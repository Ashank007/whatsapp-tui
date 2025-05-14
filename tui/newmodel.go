package tui

import (
	"whatsapp-tui/contacts"
	"go.mau.fi/whatsmeow"
)

// NewModel creates a new TUI model with the given contacts and WhatsApp client.
func NewModel(contacts []contacts.Contact, client *whatsmeow.Client) *Model {
	return &Model{
		contacts:         contacts,
		filteredContacts: contacts,
		selected:         0,
		messageLines:     []string{""},
		searchQuery:      "",
		step:             0,
		client:           client,
		statusMessage:    "Connected",
		isSending:        false,
		windowHeight:     10,
		scrollOffset:     0,
		terminalWidth:    80,
		currentLine:      0,
	}
}


