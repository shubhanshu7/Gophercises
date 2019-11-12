package cmd

import (
	"fmt"

	"github.com/shubhanshu7/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var mockShow = db.AllTasks

// CheckList will give list of all task
func CheckList(cmd *cobra.Command, args []string) {
	tasks, err := mockShow()
	if err != nil {
		fmt.Println("something is wrong", err.Error())
	}
	if len(tasks) == 0 {
		fmt.Println("nothing in list,you are lazy af!!!")
	}
	fmt.Println("you have following tasks")
	for i, task := range tasks {
		fmt.Printf("%d . %s\n", i+1, task.Task)
	}
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List of all task",
	Run:   CheckList,
}

func init() {
	RootCmd.AddCommand(listCmd)
}
