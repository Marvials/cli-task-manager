package change

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"os"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/charmbracelet/lipgloss"
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
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FB191E"))

		if strings.TrimSpace(args[1]) == "" {
			fmt.Println(style.Render("The new status cannot be empty"))
			os.Exit(1)
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

		style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3EF723"))

		fmt.Println(style.Render("Task status updated successfully!"))
	},
}

func init() {
	root.AddSubCommand(changeStatusCmd)
}
