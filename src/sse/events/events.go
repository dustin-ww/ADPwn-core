package events

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Event-Typen als Enum
const (
	EventTypeHost     = "host"
	EventTypeService  = "service"
	EventTypePassword = "password"
)

// Event-Struktur
type Event struct {
	Type    string                 `json:"type"`    // "host", "service", "password"
	ID      string                 `json:"id"`      // Eindeutige Event-ID
	Created int64                  `json:"created"` // Unix-Timestamp
	Data    map[string]interface{} `json:"data"`    // Typ-spezifische Daten
}

type Broker struct {
	clients map[chan string]bool
	add     chan chan string
	remove  chan chan string
}

func NewBroker() *Broker {
	return &Broker{
		clients: make(map[chan string]bool),
		add:     make(chan chan string),
		remove:  make(chan chan string),
	}
}

func (b *Broker) Start() {
	for {
		select {
		case client := <-b.add:
			b.clients[client] = true
		case client := <-b.remove:
			delete(b.clients, client)
			close(client)
		}
	}
}

func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// SSE Header setzen
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS

	// Client-Channel erstellen
	clientChan := make(chan string)
	b.add <- clientChan
	defer func() {
		b.remove <- clientChan
	}()

	// Verbindung offen halten
	for {
		select {
		case msg := <-clientChan:
			encoder := json.NewEncoder(w)

			for event := range clientChan {
				// SSE-Format: "event: <type>\ndata: <json>\n\n"
				fmt.Fprintf(w, "event: %s\n", event.Type)
				fmt.Fprintf(w, "data: ")
				encoder.Encode(event)
				fmt.Fprintf(w, "\n\n")
				flusher.Flush()
			}
		case <-r.Context().Done():
			return
		}
	}
}

// Funktion zum Senden von Events
func (b *Broker) SendEvent(message string) {
	for client := range b.clients {
		client <- message
	}
}
