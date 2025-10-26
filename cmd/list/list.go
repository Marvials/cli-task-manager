package list

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			log.Fatal("This command takes no arguments")
		}

		db, err := database.Connect()
		if err != nil {
			log.Fatal("Failed to connect to the database: ", err)
		}
		defer db.Close(context.Background())

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

		tasks, err := service.ListTask()
		if err != nil {
			log.Fatal("Failed to list the task: ", err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated At")

		for _, task := range tasks {

			duration := time.Since(task.CreateAt.Local())

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, task.Status, timediff.TimeDiff(time.Now().Add(-1*duration)))
		}

	},
}

func init() {
	root.AddSubCommand(listCmd)
}
