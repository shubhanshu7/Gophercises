package cmd

import (
	"fmt"
	"strings"

	// "github.com/shubhanshu7/task/db"

	"github.com/shubhanshu7/Gophercises/task/db"
	"github.com/spf13/cobra"
)

func CheckAdd(cmd *cobra.Command, args []string) {
	task := strings.Join(args, " ")
	_, err := db.CreateTask(task)
	if err != nil {
		// fmt.Print("something is wrong", err.Error())
		return
	}
	fmt.Printf("Added \"%s\" to your task list\n", task)

}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task",
	Run:   CheckAdd,
}

func init() {
	RootCmd.AddCommand(addCmd)
}
