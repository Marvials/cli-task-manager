package service

import (
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

func (s *TaskService) ListTask() ([]model.Task, error) {
	tasks, err := s.Repository.ListTodoTask()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// UpdateStatus validates the input parameters, checks if the task exists,
// and updates its status in the database. Returns an error if the ID is zero,
// the status is invalid, or the update operation fails.
func (s *TaskService) UpdateStatus(id uint, newStatus model.TaskStatus) error {
	if id == 0 {
		return errors.New("ID cannot be zero")
	}

	if !(strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DO)) ||
		strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DOING)) ||
		strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DONE))) {
		return errors.New("Status does not exist, please use one of: To do, doing, done.")
	}

	task, err := s.Repository.GetTaskByID(id)
	if err != nil {
		return err
	}

	if task == (model.Task{}) {
		return errors.New("No task exists with this ID")
	}

	var newStatusFormated model.TaskStatus

	if strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DO)) {
		newStatusFormated = model.TASK_STATUS_DO
	} else if strings.EqualFold(string(newStatus), string(model.TASK_STATUS_DOING)) {
		newStatusFormated = model.TASK_STATUS_DOING
	} else {
		newStatusFormated = model.TASK_STATUS_DONE
	}

	err = s.Repository.UpdateStatus(id, newStatusFormated)
	if err != nil {
		return err
	}

	return nil
}
