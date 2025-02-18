package logger

import (
	"sync"
)

type ADPwnLogger struct {
	subscribers []chan string
	mu          sync.Mutex
}

func NewADPwnLogger() *ADPwnLogger {
	return &ADPwnLogger{
		subscribers: make([]chan string, 0),
	}
}

func (l *ADPwnLogger) Subscribe() <-chan string {
	l.mu.Lock()
	defer l.mu.Unlock()

	ch := make(chan string, 10)
	l.subscribers = append(l.subscribers, ch)
	return ch
}

func (l *ADPwnLogger) Log(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, sub := range l.subscribers {
		sub <- message
	}
}

func (l *ADPwnLogger) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, sub := range l.subscribers {
		close(sub)
	}
}
