package whatsapp

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mau.fi/whatsmeow"
    "go.mau.fi/whatsmeow/store/sqlstore"
    waLog "go.mau.fi/whatsmeow/util/log"
    "github.com/skip2/go-qrcode"
)

// Connect initializes and connects to WhatsApp, returning a client instance.
func Connect(dbPath string) *whatsmeow.Client {
    dbLog := waLog.Stdout("DB", "INFO", true)
    container, err := sqlstore.New("sqlite3", dbPath, dbLog)
    if err != nil {
        log.Fatalf("❌ Database error: %v", err)
    }

    device, err := container.GetFirstDevice()
    if err != nil {
        log.Fatalf("❌ No device found: %v", err)
    }

    client := whatsmeow.NewClient(device, waLog.Stdout("Client", "INFO", true))

    if client.Store.ID == nil {
        qrChan, _ := client.GetQRChannel(context.Background())
        err := client.Connect()
        if err != nil {
            log.Fatalf("❌ Failed to connect: %v", err)
        }
        timeout := time.After(60 * time.Second)
        for {
            select {
            case evt, ok := <-qrChan:
                if !ok {
                    log.Fatalf("❌ QR channel closed unexpectedly")
                }
                if evt.Event == "code" {
                    qr, err := qrcode.New(evt.Code, qrcode.Medium)
                    if err != nil {
                        log.Fatalf("❌ QR code generation failed: %v", err)
                    }
                    fmt.Println("\n📲 Scan this QR code in your WhatsApp app:")
                    fmt.Println(qr.ToString(false))
                } else if evt.Event == "success" {
                    fmt.Println("✅ QR code scanned successfully!")
                    return client
                }
            case <-timeout:
                log.Fatalf("❌ QR code scanning timed out after 60 seconds")
            }
        }
    } else {
        err := client.Connect()
        if err != nil {
            log.Fatalf("❌ Connection failed: %v", err)
        }
    }

    return client
}

