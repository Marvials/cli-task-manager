package change

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/model"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/spf13/cobra"
)

var changeStatusCmd = &cobra.Command{
	Use:     "change-status [id] [new status]",
	Aliases: []string{"Change-status, CHANGE-STATUS, Change-Status"},
	Short:   "Change a task's status",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) == 1 {
			log.Fatal("Task ID and new status are required")
		}

		if strings.TrimSpace(args[1]) == "" {
			log.Fatal("The new status cannot be empty")
		}

		ctx := cmd.Context()

		db, err := database.Connect()
		if err != nil {
			log.Fatal("Failed to connect to the database: ", err)
		}
		defer db.Close(context.Background())

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Failed to convert string to int: ", err)
		}

		err = service.UpdateStatus(ctx, uint(id), model.TaskStatus(args[1]))
		if err != nil {
			log.Fatal("Failed to update task status: ", err)
		}

		log.Println("Task status updated successfully")

	},
}

func init() {
	root.AddSubCommand(changeStatusCmd)
}
