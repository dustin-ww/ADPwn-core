package sse

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// Global registry of loggers
var (
	loggers   = sync.Map{} // Stores logger instances with RunID as key
	cleanupMu sync.Mutex
)

// GetLogger returns an existing logger by runID or creates a new one
func GetLogger(runID string) *SSELogger {
	if logger, ok := loggers.Load(runID); ok {
		return logger.(*SSELogger)
	}

	logger := NewSSELogger(runID)
	loggers.Store(runID, logger)
	return logger
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters := map[string]string{
		"runId":  query.Get("runId"),
		"module": query.Get("module"),
		"level":  query.Get("level"),
		"type":   query.Get("type"),
		"since":  query.Get("since"),
	}

	logMutex.RLock()
	defer logMutex.RUnlock()

	filtered := make([]LogEntry, 0)
	for _, entry := range logStore {
		if !matchesFilters(entry, filters) {
			continue
		}
		filtered = append(filtered, entry)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func matchesFilters(entry LogEntry, filters map[string]string) bool {
	if filters["runId"] != "" && entry.RunID != filters["runId"] {
		return false
	}
	if filters["module"] != "" && entry.ModuleKey != filters["module"] {
		return false
	}
	if filters["level"] != "" && string(entry.Level) != filters["level"] {
		return false
	}
	if filters["type"] != "" && entry.EventType != filters["type"] {
		return false
	}
	if filters["since"] != "" {
		since, err := strconv.ParseInt(filters["since"], 10, 64)
		if err == nil && entry.Timestamp < since {
			return false
		}
	}
	return true
}

func CleanupOldLoggers(maxAge time.Duration) {
	loggers.Range(func(key, value interface{}) bool {
		logger := value.(*SSELogger)
		if time.Since(logger.created) > maxAge {
			loggers.Delete(key)
		}
		return true
	})
}

// GetDefaultLogger returns a system-level logger for global events
func GetDefaultLogger() *SSELogger {
	systemLoggerID := "system"
	return GetLogger(systemLoggerID)
}

// SSEHandler handles SSE connections
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Add additional CORS headers
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Get runID from query parameters to determine which logger to subscribe to
	runID := r.URL.Query().Get("runId")
	if runID == "" {
		runID = "system" // Default to system logger if no runID specified
	}

	// Get the appropriate logger
	logger := GetLogger(runID)

	// Create channel for this client
	messageChan := logger.RegisterClient()

	// Use request context for cancellation
	ctx := r.Context()

	// Clean up when connection closes
	go func() {
		<-ctx.Done() // Wait for client disconnection
		logger.UnregisterClient(messageChan)
	}()

	// Send initial connection established message
	fmt.Fprintf(w, "event: connected\n")
	fmt.Fprintf(w, "data: {\"status\":\"connected\", \"runId\":\"%s\"}\n\n", runID)
	w.(http.Flusher).Flush()

	// Event loop
	for {
		select {
		case <-ctx.Done():
			return // Client disconnected
		case event, ok := <-messageChan:
			if !ok {
				return // Channel closed
			}

			// Convert event to JSON
			jsonData, err := json.Marshal(event)
			if err != nil {
				continue
			}

			// Write in SSE format
			fmt.Fprintf(w, "event: %s\n", event.Type)
			fmt.Fprintf(w, "data: %s\n\n", jsonData)
			log.Printf("writing event: %s\n", event.Type)
			w.(http.Flusher).Flush()
		}
	}
}

// TriggerEventHandler handles HTTP requests to trigger SSE events
func TriggerEventHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for this endpoint too
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	eventType := r.URL.Query().Get("type")
	message := r.URL.Query().Get("msg")
	runID := r.URL.Query().Get("runId")

	if eventType == "" || message == "" {
		http.Error(w, "Parameters 'type' and 'msg' required", http.StatusBadRequest)
		return
	}

	// Get the appropriate loggerxx
	var logger *SSELogger
	if runID == "" {
		logger = GetDefaultLogger()
	} else {
		logger = GetLogger(runID)
	}

	// Use appropriate logger method based on eventType
	switch eventType {
	case string(DEBUG):
		logger.Debug(message)
	case string(INFO):
		logger.Info(message)
	case string(WARNING):
		logger.Warning(message)
	case string(ERROR):
		logger.Error(message)
	default:
		// Custom event type
		logger.Event(eventType, map[string]string{
			"message": message,
			"source":  "HTTP-Trigger",
		})
	}

	fmt.Fprintf(w, "Event '%s' sent", eventType)
}

// StartServer starts the SSE server on the specified port
func StartServer(port string) {
	// Initialize the system logger
	systemLogger := GetDefaultLogger()

	// Start heartbeat for system logger
	systemLogger.StartHeartbeat(5 * time.Second)

	// Log server start
	systemLogger.Info("SSE server starting", map[string]string{
		"port": port,
	})

	// Set up periodic cleanup of old loggers (every hour)
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for range ticker.C {
			CleanupOldLoggers(24 * time.Hour) // Remove loggers older than 24 hours
		}
	}()

	http.HandleFunc("/sse", SSEHandler)
	http.HandleFunc("/trigger", TriggerEventHandler)

	fmt.Printf("SSE server running on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}
