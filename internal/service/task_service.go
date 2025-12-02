package service

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/Marvials/cli-task-manager/internal/model"
	"github.com/Marvials/cli-task-manager/internal/repository"
)

type TaskService struct {
	Repository *repository.TaskRepository
}

// EnsureTaskTableExists verifies whether the task table exists in the database.
// If the table does not exist, it creates it. Returns an error if any operation fails.
func (s *TaskService) EnsureTaskTableExists(ctx context.Context) error {

	existsTable, err := s.Repository.CheckIfTaskTableExists(ctx)
	if err != nil {
		return err
	}

	if existsTable {
		return nil
	}

	err = s.Repository.CreateTaskTable(ctx)
	if err != nil {
		return err
	}

	return nil
}

// CreateTask creates a new task with the provided description.
// It initialize the task with a default status and the current timestamp,
// then persists it to the repository.
func (s *TaskService) CreateTask(ctx context.Context, description string) error {
	task := model.Task{
		Description: description,
		Status:      model.TASK_STATUS_DO,
		CreateAt:    time.Now(),
	}

	err := s.Repository.CreateTask(ctx, task)
	if err != nil {
		return err
	}

	return nil

}

// ListTasks returns tasks filtered by status.
func (s *TaskService) ListTasks(listDoingTasks, listDoneTasks, listAllTasks bool) ([]model.Task, error) {
	var err error
	var tasks []model.Task

	if listDoingTasks {
		tasks, err = s.Repository.ListDoingTasks()
	} else if listDoneTasks {
		tasks, err = s.Repository.ListDoneTasks()
	} else if listAllTasks {
		tasks, err = s.Repository.ListAllTasks()
	} else {
		tasks, err = s.Repository.ListTodoTask()
	}
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateStatus validates the input parameters, checks if the task exists,
// and updates its status in the database. Returns an error if the ID is zero,
// the status is invalid, or the update operation fails.
func (s *TaskService) UpdateStatus(ctx context.Context, id uint, newStatus model.TaskStatus) error {
	if id == 0 {
		return errors.New("ID cannot be zero")
	}

	if !(strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DO)) ||
		strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DOING)) ||
		strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DONE))) {
		return errors.New("status does not exist, please use one of: To do, doing or done")
	}

	var newStatusFormatted model.TaskStatus

	if strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DO)) {
		newStatusFormatted = model.TASK_STATUS_DO
	} else if strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DOING)) {
		newStatusFormatted = model.TASK_STATUS_DOING
	} else {
		newStatusFormatted = model.TASK_STATUS_DONE
	}

	err := s.Repository.UpdateStatus(ctx, id, newStatusFormatted)
	if err != nil {
		return err
	}

	return nil
}

// GetTask retrieves a specific task by its unique identifier (ID).
// It performs a validation to ensure the ID is non-zero before querying the repository.
func (s *TaskService) GetTask(ctx context.Context, id uint) (model.Task, error) {
	if id == 0 {
		return model.Task{}, errors.New("ID cannot be zero")
	}

	task, err := s.Repository.GetTaskByID(ctx, id)
	if err != nil {
		return model.Task{}, err
	}

	if task == (model.Task{}) {
		return model.Task{}, errors.New("no task exits with this ID")
	}

	return task, nil
}

// DeleteTask removes a task from the database by its ID.
func (s *TaskService) DeleteTask(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("ID cannot be zero")
	}

	err := s.Repository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil

}
