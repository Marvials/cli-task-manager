package get

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Show details of a specific task",
	Long: `Retrieves a task by its unique ID and displays all information,
including how long aog it was created`,
	Example: "task get 4",
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"Get", "GET"},
	Run: func(cmd *cobra.Command, args []string) {
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FB191E"))

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(style.Render("Invalid ID format. ID must be an integer"))
			os.Exit(1)
		}

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close(context.Background())

		task, err := service.GetTask(ctx, uint(id))
		if err != nil {
			log.Fatal(err)
		}

		var (
			blue      = lipgloss.Color("#3462FA")
			lightGray = lipgloss.Color("#7B7F84")

			headerStyle  = lipgloss.NewStyle().Foreground(blue).Bold(true).Align(lipgloss.Center)
			cellStyle    = lipgloss.NewStyle().Padding(0, 1).Width(30)
			rowStyle = cellStyle.Foreground(lightGray)
		)

		t := table.New().
			Border(lipgloss.NormalBorder()).
			BorderStyle(lipgloss.NewStyle().Foreground(blue)).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch row {
				case table.HeaderRow:
					return headerStyle
				default:
					return rowStyle
				}
			}).Headers("ID", "DESCRIPTION", "STATUS", "CREATED AT")

		duration := time.Since(task.CreateAt.Local())
		timediff := timediff.TimeDiff(time.Now().Add(-1 * duration))
		IDString := fmt.Sprint(task.ID)

		t.Row(IDString, task.Description, string(task.Status), timediff)

		fmt.Println(t)
	},
}

func init() {
	root.AddSubCommand(getCmd)
}
