package root

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI Task Manager",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing the CLI Task manager: ", err)
	}
}
