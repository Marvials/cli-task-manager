package get

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/database"
	"github.com/Marvials/cli-task-manager/internal/repository"
	"github.com/Marvials/cli-task-manager/internal/service"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [id]",
	Short:   "Get task by ID",
	Aliases: []string{"Get", "GET"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatal("Task ID is required")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Invalid ID format. ID must be an integer")
		}

		db, err := database.Connect()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close(context.Background())

		repo := repository.NewTaskRepository(db)
		service := service.TaskService{Repository: repo}

		task, err := service.GetTask(uint(id))
		if err != nil {
			log.Fatal(err)
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 10, ' ', 0)
		defer w.Flush()

		fmt.Fprintln(w, "ID\tDescription\tStatus\tCreated At")

		duration := time.Since(task.CreateAt.Local())
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Description, task.Status, timediff.TimeDiff(time.Now().Add(-1*duration)))

	},
}

func init() {
	root.AddSubCommand(getCmd)
}
