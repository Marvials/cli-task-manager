package add

import (
	"context"
	"log"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
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

		description := strings.Join(args, " ")

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		err = service.CreateTask(ctx, description)
		if err != nil {
			log.Fatal("Error creating the task: ", err)
		}

		log.Println("Task was created successfully!")
	},
}

func init() {
	root.AddSubCommand(addCmd)
}
