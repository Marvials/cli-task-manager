package model

import (
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_DO    TaskStatus = "To do"
	TASK_STATUS_DOING TaskStatus = "Doing"
	TASK_STATUS_DONE  TaskStatus = "Done"
)

type Task struct {
	ID          uint
	Description string
	Status      TaskStatus
	CreateAt    time.Time
}
