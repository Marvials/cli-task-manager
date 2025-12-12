package list

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing tasks",
	Long: `Display a table with your tasks.
By default, it shows only pending ('To do') tasks.
Use available flags to filter by other statuses`,
	Example: `task list   # List tasks to do
task list --doing  #List tasks in progress
task list --done  #List completed tasks
task list --all  #List everything`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
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

		var (
			blue = lipgloss.Color("#3462FA")
			gray = lipgloss.Color("245")
			lightGray = lipgloss.Color("241")

			headerStyle = lipgloss.NewStyle().Foreground(blue).Bold(true).Align(lipgloss.Center)
			cellStyle = lipgloss.NewStyle().Padding(0, 1).Width(30)
			oddRowStyle = cellStyle.Foreground(gray)
			evenRowStyle = cellStyle.Foreground(lightGray)
		)


		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(blue)).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == table.HeaderRow:
					return headerStyle
				case row%2 == 0:
					return evenRowStyle
				default:
					return oddRowStyle
				}
			}).Headers("ID", "DESCRIPTION", "STATUS", "CREATED AT")

		for _, task := range tasks {
			duration := time.Since(task.CreateAt.Local())
			timediff := timediff.TimeDiff(time.Now().Add(-1 * duration))
			IDString := fmt.Sprint(task.ID)

			t.Row(IDString, task.Description, string(task.Status), timediff)
		}

		fmt.Println(t)

	},
}

func init() {
	listCmd.Flags().Bool("done", false, "List all tasks with status in done")
	listCmd.Flags().Bool("doing", false, "List all tasks with status in doing")
	listCmd.Flags().Bool("all", false, "List all tasks")

	listCmd.MarkFlagsMutuallyExclusive("done", "doing", "all")

	root.AddSubCommand(listCmd)
}
