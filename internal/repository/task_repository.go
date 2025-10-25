package repository

import (
	"context"
	"time"

	"github.com/Marvials/cli-task-manager/internal/model"
	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	db *pgx.Conn
}

// NewTaskRepository returns a new instance of TaskRepository
// with an active database connection.
func NewTaskRepository(conn *pgx.Conn) *TaskRepository {
	return &TaskRepository{conn}
}

// CheckIfTaskTableExists verifies whether the tasks table exists in the database.
// It returns true if the table exists, otherwise false.
func (r *TaskRepository) CheckIfTaskTableExists() (bool, error) {
	query := `SELECT to_regclass('public.tasks') IS NOT NULL;`

	var existsTable bool
	if err := r.db.QueryRow(context.Background(), query).Scan(&existsTable); err != nil {
		return false, err
	}

	return existsTable, nil
}

// CreateTaskTable creates the tasks table in the database if it does not already exist.
func (r *TaskRepository) CreateTaskTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			description VARCHAR(100) NOT NULL,
			status VARCHAR(5) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	ctx := context.Background()
	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

// CreateTask adds a task to the tasks table in the database
func (r *TaskRepository) CreateTask(task model.Task) error {
	query := `
		INSERT INTO tasks (description, status, created_at) VALUES ($1, $2, $3) RETURNING ID
	`

	var id int
	err := r.db.QueryRow(context.Background(), query, task.Description, task.Status, task.CreateAt).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

// ListTodoTask retriaves tasks with "to do" status from the database
func (r *TaskRepository) ListTodoTask() ([]model.Task, error) {
	query := `
		SELECT id, description, status, created_at FROM tasks
		WHERE status = $1;
	`

	rows, err := r.db.Query(context.Background(), query, model.TASK_STATUS_DO)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var id uint
		var description, status string
		var createAt time.Time

		err := rows.Scan(&id, &description, &status, &createAt)
		if err != nil {
			return nil, err
		}

		task := model.Task{
			ID:          id,
			Description: description,
			Status:      model.TaskStatus(status),
			CreateAt:    createAt,
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}
