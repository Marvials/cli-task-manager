package list

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
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

		listDoingTasks, err := cmd.Flags().GetBool("doing")
		if err != nil {
			log.Fatal("Failed to get the value of flag doing: ", err)
		}
		listDoneTasks, err := cmd.Flags().GetBool("done")
		if err != nil {
			log.Fatal("Failed to get the value of flag done: ", err)
		}
		listAllTasks, err := cmd.Flags().GetBool("all")
		if err != nil {
			log.Fatal("Failed to get the value of flag all: ", err)
		}

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		tasks, err := service.ListTasks(ctx, listDoingTasks, listDoneTasks, listAllTasks)
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
	listCmd.Flags().Bool("done", false, "List all tasks with status in done")
	listCmd.Flags().Bool("doing", false, "List all tasks with status in doing")
	listCmd.Flags().Bool("all", false, "List all tasks")

	listCmd.MarkFlagsMutuallyExclusive("done", "doing", "all")

	root.AddSubCommand(listCmd)
}
