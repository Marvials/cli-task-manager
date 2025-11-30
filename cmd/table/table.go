package table

import (
	"context"
	"log"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "tables",
	Short: "Verify and create tables in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			log.Fatal("This command takes no arguments")
		}

		ctx := cmd.Context()

		db, err := database.Connect()
		if err != nil {
			log.Fatal("Error connecting to the database: ", err)
		}
		defer db.Close(context.Background())

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

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
