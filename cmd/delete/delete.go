package delete

import (
	"log"
	"strconv"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Short:   "Delete task by ID",
	Aliases: []string{"DELETE", "Delete"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Task id must be a number")
		}

		db, err := database.Connect()
		if err != nil {
			log.Fatal("Failed to connect to the database: ", err)
		}

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

		err = service.DeleteTask(uint(id))
		if err != nil {
			log.Fatal("Failed to delete the task: ", err)
		}
	},
}

func init() {
	root.AddSubCommand(deleteCmd)
}
