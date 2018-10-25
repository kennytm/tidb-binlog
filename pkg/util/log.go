package util

import (
	"sync"
	"time"
)

// Log prints log only after a certain amount of time
type Log struct {
	sync.RWMutex
	lastTime map[string]time.Time
	interval map[string]time.Duration
}

// NewLog returns a new Log
func NewLog() *Log {
	return &Log{
		lastTime: make(map[string]time.Time),
		interval: make(map[string]time.Duration),
	}
}

// Add adds new label
func (l *Log) Add(label string, interval time.Duration) {
	l.Lock()
	l.interval[label] = interval
	l.Unlock()
}

// Print executes the fn to print log
func (l *Log) Print(label string, fn func()) {
	l.Lock()
	defer l.Unlock()

	_, ok := l.lastTime[label]
	if !ok || time.Since(l.lastTime[label]) > l.interval[label] {
		fn()
		l.lastTime[label] = time.Now()
	}
}
