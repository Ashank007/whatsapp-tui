
package main

import (
    "fmt"
    "os"

    bubbletea "github.com/charmbracelet/bubbletea"
    _ "github.com/mattn/go-sqlite3"

    "whatsapp-tui/contacts"
    "whatsapp-tui/tui"
    "whatsapp-tui/whatsapp"
)

func main() {
    // Load contacts
    contactsList := contacts.LoadContacts("contacts.json")

    // Connect to WhatsApp
    client := whatsapp.Connect("file:session.db?_foreign_keys=on")

    // Start the TUI
    p := bubbletea.NewProgram(tui.NewModel(contactsList, client), bubbletea.WithAltScreen())

    if err := p.Start(); err != nil {
        fmt.Println("💥 App crashed:", err)
        os.Exit(1)
    }
}

