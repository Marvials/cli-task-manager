package table

import (
	"context"
	"log"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "tables",
	Short: "Verify and create tables in the database",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		err = service.EnsureTaskTableExists(ctx)
		if err != nil {
			log.Fatal("Error creating the task table: ", err)
		}

		log.Println("Task table created successfully")
	},
}

func init() {
	root.AddSubCommand(tableCmd)
}
