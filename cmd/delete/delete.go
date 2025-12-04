package delete

import (
	"context"
	"log"
	"strconv"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
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

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		err = service.DeleteTask(ctx, uint(id))
		if err != nil {
			log.Fatal("Failed to delete the task: ", err)
		}
	},
}

func init() {
	root.AddSubCommand(deleteCmd)
}
