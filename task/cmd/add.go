package cmd

import(
	"fmt"
	"strings"

	"github.com/dlokkers/gophercises/task/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task to the TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if err := tasklist.Add(task); err != nil{
			fmt.Println("Something went wrong: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
