package sse

import (
	"log"
	"strings"
	"sync"
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel string

const (
	DEBUG   LogLevel = "debug"
	INFO    LogLevel = "info"
	WARNING LogLevel = "warning"
	ERROR   LogLevel = "error"
)

var (
	logStore  []LogEntry
	logMutex  sync.RWMutex
	maxLogs   = 10000
	pruneSize = 2000
)

// LogEvent represents a log message to be sent to the frontend
type LogEvent struct {
	RunID     string      `json:"runId"`
	Level     LogLevel    `json:"level"`
	Message   string      `json:"message"`
	Details   interface{} `json:"details,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

type LogEntry struct {
	RunID     string      `json:"runId"`
	ModuleKey string      `json:"moduleKey,omitempty"`
	Level     LogLevel    `json:"level,omitempty"`
	EventType string      `json:"eventType,omitempty"`
	Message   string      `json:"message,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

// SSEEvent represents an event to be sent over Server-Sent Events
type SSEEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
	ID      int         `json:"id"`
	RunID   string      `json:"runId"` // RunID im SSEEvent
}

// SSELogger provides a clean interface for sending log events to connected clients
type SSELogger struct {
	clients   map[chan SSEEvent]bool
	mutex     sync.Mutex
	eventID   int
	eventBus  chan SSEEvent
	runID     string
	moduleKey string
	created   time.Time
}

// NewSSELogger creates a new SSELogger instance
func NewSSELogger(runID string) *SSELogger {
	logger := &SSELogger{
		clients:  make(map[chan SSEEvent]bool),
		eventBus: make(chan SSEEvent, 100),
		runID:    runID,
		created:  time.Now(),
	}
	go logger.dispatchEvents()
	return logger
}

func (l *SSELogger) ForModule(moduleKey string) *SSELogger {
	return &SSELogger{
		runID:     l.runID,
		clients:   l.clients,
		eventBus:  l.eventBus,
		moduleKey: moduleKey,
		created:   time.Now(),
	}
}

// RegisterClient registers a new SSE client channel
func (l *SSELogger) RegisterClient() chan SSEEvent {
	clientChan := make(chan SSEEvent, 10) // Buffer to avoid blocking

	l.mutex.Lock()
	l.clients[clientChan] = true
	l.mutex.Unlock()

	return clientChan
}

// UnregisterClient removes a client channel
func (l *SSELogger) UnregisterClient(clientChan chan SSEEvent) {
	l.mutex.Lock()
	delete(l.clients, clientChan)
	close(clientChan)
	l.mutex.Unlock()
}

func (l *SSELogger) Log(level LogLevel, message string, details interface{}) {
	entry := LogEntry{
		RunID:     l.runID,
		ModuleKey: l.moduleKey,
		Level:     level,
		Message:   message,
		Payload:   details,
		Timestamp: time.Now().Unix(),
	}

	l.storeLog(entry)
	l.sendSSEEvent("log", entry)
	// The system logger forwarding is now handled in sendSSEEvent
}

func (l *SSELogger) Event(eventType string, payload interface{}) {
	entry := LogEntry{
		RunID:     l.runID,
		ModuleKey: l.moduleKey,
		EventType: eventType,
		Payload:   payload,
		Timestamp: time.Now().Unix(),
	}

	l.storeLog(entry)
	l.sendSSEEvent(eventType, entry)
	// The system logger forwarding is now handled in sendSSEEvent
}

func (l *SSELogger) storeLog(entry LogEntry) {
	logMutex.Lock()
	defer logMutex.Unlock()

	if len(logStore) >= maxLogs {
		logStore = logStore[pruneSize:]
	}
	// Fix the printf formatting
	log.Printf("Storing new log: %s", entry.Message)
	logStore = append(logStore, entry)
}

// In your sse package, add a function to forward events to system logger
func (l *SSELogger) sendToSystemLogger(eventType string, payload interface{}) {
	// Only forward if this isn't already the system logger
	if l.runID != "system" {
		sysLogger := GetDefaultLogger()
		// Add the original runID to the payload if it's a map
		if payloadMap, ok := payload.(map[string]interface{}); ok {
			if _, exists := payloadMap["originalRunId"]; !exists {
				payloadMap["originalRunId"] = l.runID
			}
		}
		sysLogger.Event(eventType, payload)
	}
}

// Modify the sendSSEEvent method to also send to system logger
func (l *SSELogger) sendSSEEvent(eventType string, payload interface{}) {
	l.mutex.Lock()
	l.eventID++
	eventID := l.eventID
	l.mutex.Unlock()

	event := SSEEvent{
		Type:    eventType,
		Payload: payload,
		ID:      eventID,
		RunID:   l.runID,
	}

	l.eventBus <- event

	// Also send module-related events to the system logger
	if eventType == "log" || strings.HasPrefix(eventType, "module_") ||
		eventType == "run_start" || eventType == "run_complete" || eventType == "run_error" {
		l.sendToSystemLogger(eventType, payload)
	}
}

// dispatchEvents processes events from the eventBus and distributes them to clients
func (l *SSELogger) dispatchEvents() {
	for event := range l.eventBus {
		l.mutex.Lock()
		for clientChan := range l.clients {
			select {
			case clientChan <- event:
				// Message sent successfully
			default:
				// Channel is full or blocked, skip this client
			}
		}
		l.mutex.Unlock()
	}
}

// GetRunID returns the runID associated with this logger
func (l *SSELogger) GetRunID() string {
	return l.runID
}

// Debug logs a debug message
func (l *SSELogger) Debug(message string, details ...interface{}) {
	var detailsData interface{}
	if len(details) == 1 {
		detailsData = details[0]
	} else if len(details) > 1 {
		detailsData = details
	}
	l.Log(DEBUG, message, detailsData)
}

// Info logs an info message
func (l *SSELogger) Info(message string, details ...interface{}) {
	var detailsData interface{}
	if len(details) == 1 {
		detailsData = details[0]
	} else if len(details) > 1 {
		detailsData = details
	}
	l.Log(INFO, message, detailsData)
}

// Warning logs a warning message
func (l *SSELogger) Warning(message string, details ...interface{}) {
	var detailsData interface{}
	if len(details) == 1 {
		detailsData = details[0]
	} else if len(details) > 1 {
		detailsData = details
	}
	l.Log(WARNING, message, detailsData)
}

// Error logs an error message
func (l *SSELogger) Error(message string, details ...interface{}) {
	var detailsData interface{}
	if len(details) == 1 {
		detailsData = details[0]
	} else if len(details) > 1 {
		detailsData = details
	}
	l.Log(ERROR, message, detailsData)
}

// StartHeartbeat begins sending periodic heartbeat events
func (l *SSELogger) StartHeartbeat(interval time.Duration) {
	/*go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			l.Event("heartbeat", map[string]interface{}{
				"timestamp": time.Now().Unix(),
				"status":    "OK",
			})
		}
	}()*/
}
