package cmd

import (
	"fmt"

	// "github.com/shubhanshu7/task/db"
	"github.com/shubhanshu7/Gophercises/task/db"
	"github.com/spf13/cobra"
)

func CheckList(cmd *cobra.Command, args []string) {
	tasks, err := db.AllTasks()
	if err != nil {
		fmt.Println("something is wrong", err.Error())
	}
	if len(tasks) == 0 {
		fmt.Println("nothing in list,you are lazy af!!!")
	}
	fmt.Println("you have following tasks")
	for i, task := range tasks {
		fmt.Printf("%d . %s\n", i+1, task.Value)
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
