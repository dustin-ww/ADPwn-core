// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type SSEEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	ID      int         `json:"id"`
}

// Globale Variablen für Client-Verwaltung
var (
	clients = make(map[chan SSEEvent]bool)
	mutex   sync.Mutex
	eventID = 0
)

func sseHandler(w http.ResponseWriter, r *http.Request) {
	// SSE-Header setzen
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Channel für diesen Client erstellen
	messageChan := make(chan SSEEvent)

	// Client registrieren
	mutex.Lock()
	clients[messageChan] = true
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, messageChan)
		mutex.Unlock()
		close(messageChan)
	}()

	// Close Notifier für Verbindungsstatus
	notifier := w.(http.CloseNotifier).CloseNotify()

	for {
		select {
		case <-notifier:
			return // Client hat Verbindung getrennt
		case event := <-messageChan:
			// Event in JSON umwandeln
			jsonData, _ := json.Marshal(event)

			// SSE-Format schreiben
			fmt.Fprintf(w, "event: %s\n", event.Type)
			fmt.Fprintf(w, "data: %s\n\n", jsonData)

			// Flushen damit Client sofort erhält
			w.(http.Flusher).Flush()
		}
	}
}

// Hilfsfunktion zum Senden von Events an alle Clients
func sendEventToAll(eventType string, payload interface{}) {
	mutex.Lock()
	defer mutex.Unlock()

	eventID++
	event := SSEEvent{
		Type:    eventType,
		Payload: payload,
		ID:      eventID,
	}

	for clientChan := range clients {
		clientChan <- event
	}
}

// Beispiel: HTTP-Endpoint zum Auslösen von Events
func triggerEventHandler(w http.ResponseWriter, r *http.Request) {
	eventType := r.URL.Query().Get("type")
	message := r.URL.Query().Get("msg")

	if eventType == "" || message == "" {
		http.Error(w, "Parameter 'type' und 'msg' benötigt", http.StatusBadRequest)
		return
	}

	sendEventToAll(eventType, map[string]string{
		"message": message,
		"source":  "HTTP-Trigger",
	})

	fmt.Fprintf(w, "Event '%s' gesendet", eventType)
}

func main() {
	http.HandleFunc("/sse", sseHandler)
	http.HandleFunc("/trigger", triggerEventHandler)

	// Beispiel: Periodische Events senden
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-ticker.C:
				sendEventToAll("heartbeat", map[string]interface{}{
					"timestamp": time.Now().Unix(),
					"status":    "OK",
				})
			}
		}
	}()

	fmt.Println("SSE-Server läuft auf :8082")
	http.ListenAndServe(":8082", nil)
}
