package factory

import (
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/jackc/pgx/v5"
)

// NewTaskService initializes the database connection and wires up the repository and service layers.
// It returns the active database connection (to allow deferred closing by the caller) and
// a fully configured instance of TaskService.
func NewTaskService() (*pgx.Conn, *service.TaskService, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, nil, err
	}

	repo := repository.NewTaskRepository(db)
	service := &service.TaskService{Repository: repo}

	return db, service, nil
}
