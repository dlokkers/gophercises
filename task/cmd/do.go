package cmd

import(
	"fmt"
	"strings"

	"github.com/dlokkers/gophercises/task/tasks"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use: "do",
	Short: "Marks a task as done on the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		key := strings.Join(args, " ")
		if err := tasklist.Do(key); err != nil {
			fmt.Println("Something went wrong:, ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
