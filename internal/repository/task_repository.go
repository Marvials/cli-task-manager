package repository

import (
	"context"
	"errors"
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
func (r *TaskRepository) CheckIfTaskTableExists(ctx context.Context) (bool, error) {
	query := `SELECT to_regclass('public.tasks') IS NOT NULL;`

	var existsTable bool
	if err := r.db.QueryRow(ctx, query).Scan(&existsTable); err != nil {
		return false, err
	}

	return existsTable, nil
}

// CreateTaskTable creates the tasks table in the database if it does not already exist.
func (r *TaskRepository) CreateTaskTable(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			description VARCHAR(100) NOT NULL,
			status VARCHAR(5) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := r.db.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

// CreateTask adds a task to the tasks table in the database
func (r *TaskRepository) CreateTask(ctx context.Context, task model.Task) error {
	query := `
		INSERT INTO tasks (description, status, created_at) VALUES ($1, $2, $3) RETURNING ID
	`

	var id int
	err := r.db.QueryRow(ctx, query, task.Description, task.Status, task.CreateAt).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

// ListTodoTask retrieves tasks with "to do" status from the database
func (r *TaskRepository) ListTodoTask() ([]model.Task, error) {
	query := `
		SELECT id,
		description,
		status,
		created_at AT TIME ZONE 'America/Sao_Paulo' AS created_at_local
		FROM tasks
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

// ListDoingTasks retrieves all tasks currently in the "DOING" status from the database.
// It returns a slice of tasks with their creation time adjusted to the 'America/Sao_Paulo' timezone,
// or an error if the query fails.
func (r *TaskRepository) ListDoingTasks() ([]model.Task, error) {
	query := `
		SELECT id, description, status, created_at AT TIME ZONE 'America/Sao_Paulo' AS created_at_local
		FROM tasks
		WHERE status = $1;
	`

	rows, err := r.db.Query(context.Background(), query, model.TASK_STATUS_DOING)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var id uint
		var description, status string
		var createdAt time.Time

		err = rows.Scan(&id, &description, &status, &createdAt)
		if err != nil {
			return nil, err
		}

		task := model.Task{
			ID:          id,
			Description: description,
			Status:      model.TaskStatus(status),
			CreateAt:    createdAt,
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

// ListDoneTasks returns all tasks with status DONE.
// Executes a query in the database filtering by status, scans each row,
// and builds a slice of model.Task for return.
func (r *TaskRepository) ListDoneTasks() ([]model.Task, error) {
	query := `
		SELECT id, description, status, created_at AT TIME ZONE 'America/Sao_Paulo' AS created_at_local
		FROM tasks
		WHERE status = $1;
	`

	rows, err := r.db.Query(context.Background(), query, model.TASK_STATUS_DONE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var id uint
		var description, status string
		var createdAt time.Time

		err = rows.Scan(&id, &description, &status, &createdAt)
		if err != nil {
			return nil, err
		}

		task := model.Task{
			ID:          id,
			Description: description,
			Status:      model.TaskStatus(status),
			CreateAt:    createdAt,
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

// ListAllTasks retrieves all tasks from the database.
func (r *TaskRepository) ListAllTasks() ([]model.Task, error) {
	query := `
		SELECT id, description, status, created_at AT TIME ZONE 'America/Sao_Paulo' AS created_at_local
		FROM tasks
		ORDER BY id
	`
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var id uint
		var description, status string
		var createdAt time.Time

		err = rows.Scan(&id, &description, &status, &createdAt)
		if err != nil {
			return nil, err
		}

		task := model.Task{
			ID:          id,
			Description: description,
			Status:      model.TaskStatus(status),
			CreateAt:    createdAt,
		}

		tasks = append(tasks, task)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil
}

// GetTaskByID retrieves a task from the database by its ID.
// It returns the task if found, or an empty Task and an error if not found or if a query error occurs.
func (r *TaskRepository) GetTaskByID(id uint) (model.Task, error) {
	query := `
		SELECT id, description, status, created_at AT TIME ZONE 'America/Sao_Paulo' AS created_at_local
		FROM tasks
		WHERE id = $1
	`

	row, err := r.db.Query(context.Background(), query, id)
	if err != nil {
		return model.Task{}, err
	}
	defer row.Close()

	var task model.Task

	if row.Next() {
		err = row.Scan(&task.ID, &task.Description, &task.Status, &task.CreateAt)
		if err != nil {
			return model.Task{}, nil
		}
	}

	if row.Err() != nil {
		return model.Task{}, err
	}

	return task, nil

}

// UpdateStatus updates the status of a task identified by its ID in the database.
// Returns an error if the update fails or if no records are affected.
func (r *TaskRepository) UpdateStatus(id uint, newStatus model.TaskStatus) error {
	query := `
		UPDATE tasks SET status = $1 WHERE id = $2;
	`

	cmdTag, err := r.db.Exec(context.Background(), query, newStatus, id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("no records were updated")
	}

	return nil
}

// DeleteTask removes a task from the database identified by its ID.
//
// It returns an error if the database operation fails or if no record was
// found to be deleted.
func (r *TaskRepository) DeleteTask(ID uint) error {
	query := `
		DELETE FROM tasks WHERE ID = $1
	`

	cmdTag, err := r.db.Exec(context.Background(), query, ID)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("no records were deleted")
	}

	return nil
}
