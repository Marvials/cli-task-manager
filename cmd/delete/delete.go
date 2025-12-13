package delete

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Remove a task permanently",
	Long: `Deletes a task from the database using its ID.
WARNING: This action cannot be undone`,
	Example: "task delete 10",
	Aliases: []string{"DELETE", "Delete"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		style := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FB191E"))

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(style.Render("Task id must be a number"))
			os.Exit(1)
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

		style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3EF723"))
		fmt.Println(style.Render("Task deleted successfully!"))
	},
}

func init() {
	root.AddSubCommand(deleteCmd)
}
