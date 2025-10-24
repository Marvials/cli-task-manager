package service

import (
	"time"

	"github.com/Marvials/cli-task-manager/internal/model"
	"github.com/Marvials/cli-task-manager/internal/repository"
)

type TaskService struct {
	Repository *repository.TaskRepository
}

// EnsureTaskTableExists verifies whether the task table exists in the database.
// If the table does not exist, it creates it. Returns an error if any operation fails.
func (s *TaskService) EnsureTaskTableExists() error {

	existsTable, err := s.Repository.CheckIfTaskTableExists()
	if err != nil {
		return err
	}

	if existsTable {
		return nil
	}

	err = s.Repository.CreateTaskTable()
	if err != nil {
		return err
	}

	return nil
}

func (s *TaskService) CreateTask(description string) error {
	task := model.Task{
		Description: description,
		Status:      model.TASK_STATUS_DO,
		CreateAt:    time.Now(),
	}

	err := s.Repository.CreateTask(task)
	if err != nil {
		return err
	}

	return nil

}
