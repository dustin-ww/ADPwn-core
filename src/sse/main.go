package main

import (
	"ADPwn/sse/events"
	"fmt"
	"net/http"
	"time"
)

// main.go
func main() {
	broker := events.NewBroker()
	go broker.Start()

	http.Handle("/events", broker)

	// Simuliere Events (z. B. aus anderem Code)
	go func() {
		for {
			time.Sleep(2 * time.Second)
			broker.SendEvent(fmt.Sprintf("Event: %s", time.Now().String()))
		}
	}()

	http.ListenAndServe(":6001", nil)
}
