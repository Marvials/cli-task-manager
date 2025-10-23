package add

import (
	"log"
	"strings"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a task in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("The description is required")
		}

		if strings.TrimSpace(args[0]) == "" {
			log.Fatal("The description cannot be empty")
		}

		log.Println("The task was created successfully!")
	},
}

func init() {
	root.AddSubCommand(addCmd)
}
