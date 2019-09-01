package cmd

import(
	"github.com/dlokkers/gophercises/task/tasks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists all tasks in the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		tasklist.List()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
