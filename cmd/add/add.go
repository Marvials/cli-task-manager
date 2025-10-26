package add

import (
	"context"
	"log"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [description]",
	Aliases: []string{"Add", "ADD"},
	Short:   "Add a task in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("The description is required")
		}

		if strings.TrimSpace(args[0]) == "" {
			log.Fatal("The description cannot be empty")
		}

		if len(args) > 1 {
			log.Fatal(`The description must be surrounded with ""`)
		}

		db, err := database.Connect()
		if err != nil {
			log.Fatal("Failed to connect to the database: ", err)
		}
		defer db.Close(context.Background())

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

		err = service.CreateTask(args[0])
		if err != nil {
			log.Fatal("Error creating the task: ", err)
		}

		log.Println("Task was created successfully!")
	},
}

func init() {
	root.AddSubCommand(addCmd)
}
