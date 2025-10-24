package repository

import (
	"context"

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
