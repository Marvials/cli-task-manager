package model

import (
	"fmt"
	"strings"
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

func ParseTaskStatus(statusRaw string) (TaskStatus, error) {
	switch {
	case strings.EqualFold(statusRaw, string(TASK_STATUS_DO)):
		return TASK_STATUS_DO, nil
	case strings.EqualFold(statusRaw, string(TASK_STATUS_DOING)):
		return TASK_STATUS_DOING, nil
	case strings.EqualFold(statusRaw, string(TASK_STATUS_DONE)):
		return TASK_STATUS_DONE, nil
	default:
		return "", fmt.Errorf("status does not exist, please use one of: To do, doing or done")
	}
}
