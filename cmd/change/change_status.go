package change

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/spf13/cobra"
)

var changeStatusCmd = &cobra.Command{
	Use:     "change-status [id] [new status]",
	Aliases: []string{"change", "update"},
	Short:   "Update a task's status",
	Long: `Changes the status of an existing task based on its ID.
Accepted statures are:
- To do
- Doing
- Done`,
	Example: `task change 1 Doing
task update 5 done`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(args[1]) == "" {
			log.Fatal("The new status cannot be empty")
		}

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Failed to convert string to int: ", err)
		}

		err = service.UpdateStatus(ctx, uint(id), args[1])
		if err != nil {
			log.Fatal("Failed to update task status: ", err)
		}

		log.Println("Task status updated successfully")

	},
}

func init() {
	root.AddSubCommand(changeStatusCmd)
}
