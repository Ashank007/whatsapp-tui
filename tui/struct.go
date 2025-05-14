package tui

import (
	"whatsapp-tui/contacts"
	"go.mau.fi/whatsmeow"

)

// Model represents the state of the TUI application.
type Model struct {
	contacts         []contacts.Contact
	filteredContacts []contacts.Contact
	selected         int
	messageLines     []string
	searchQuery      string
	step             int
	client           *whatsmeow.Client
	statusMessage    string
	isSending        bool
	windowHeight     int
	scrollOffset     int
	terminalWidth    int
	currentLine      int
	cursorPos        int
}
