package models

import (
	"fmt"
	"time"
)

var Storage = map[string]*Task{
	"1": {
		Title:   "1",
		Status:  "In Progress",
		Created: time.Now(),
	},
}

type Task struct {
	Title   string    `json:"title"`
	Status  string    `json:"status"`
	Created time.Time `json:"created"`
}

type TasksAutoProcess interface {
	LaunchAutoProcess()
}

func (t *Task) LaunchAutoProcess() {
	t.Status = "In Progress"
	go func(status string) {
		defer fmt.Println("end")
		status = "In Progress"
		time.Sleep(1 * time.Minute)
		status = "Done"
	}(t.Status)
}
