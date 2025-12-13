package add

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/Marvials/cli-task-manager/internal/factory"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [description]",
	Aliases: []string{"Add", "ADD"},
	Short:   "Create a new task",
	Long: `Adds a new task to your 'To Do' list.
You can type the description as a normal sentence without needing quotes.`,
	Example: `task add Buy milk
task add "Review project code"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FB191E"))
		if strings.TrimSpace(args[0]) == "" {
			fmt.Println(style.Render("The description cannot be empty"))
			os.Exit(1)
		}

		description := strings.Join(args, " ")

		ctx := cmd.Context()

		db, service, err := factory.NewTaskService()
		if err != nil {
			log.Fatal("Failed to initialize dependencies: ", err)
		}
		defer db.Close(context.Background())

		err = service.CreateTask(ctx, description)
		if err != nil {
			log.Fatal("Error creating the task: ", err)
		}

		style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#3EF723"))

		fmt.Println(style.Render("Task was created successfully!"))
	},
}

func init() {
	root.AddSubCommand(addCmd)
}
