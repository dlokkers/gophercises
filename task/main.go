package main

import(
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/dlokkers/gophercises/task/cmd"
	"github.com/dlokkers/gophercises/task/tasks"
)

func main() {
	home, _ := homedir.Dir()
	tasklist.Init(filepath.Join(home, ".tasks"))
	cmd.Execute()
}
