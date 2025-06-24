package models

import (
	"sync"
	"time"
)

type Task struct {
	Title   string    `json:"title"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
	mu      sync.RWMutex
}

func (t *Task) LaunchAutoProcess() {
	go func() {
		t.mu.Lock()
		defer t.mu.Unlock()
		t.Status = "In Progress"
		time.Sleep(3 * time.Minute)
		t.Status = "Done"
	}()
}
