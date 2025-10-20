package table

import (
	"log"

	"github.com/Marvials/cli-task-manager/cmd/root"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "tables",
	Short: "Verify and create tables in the database",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			log.Fatal("This command does not have arguments")
		}

		log.Println("The table has been created")
	},
}

func init() {
	root.AddSubCommand(tableCmd)
}
