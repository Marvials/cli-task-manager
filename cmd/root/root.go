package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A simple and efficient CLI task manager",
	Long: `Task Manager is a command-line tool to manage your day-to-day tasks.
	It allows you to create, list, update, and remove tasks quickly, keeping
	everything organized in a PostgresSQL database.`,
}

// Execute runs the root command of the CLI application.
// It initializes the command hierarchy and handles any execution errors.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing the CLI Task manager: ", err)
	}
}

// AddSubCommand adds a subcomamand to the root command
func AddSubCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}
