package cmd

import (
	"fmt"
	"strconv"

	"github.com/shubhanshu7/Gophercises/task/db"
	"github.com/spf13/cobra"
)

var mockRemove = db.DeleteTask

// Delete Will delete task
func Delete(cmd *cobra.Command, args []string) {
	var ids []int
	for _, arg := range args {
		id, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("failed to parse", arg)
		} else {
			ids = append(ids, id)
		}
	}
	tasks, err := mockShow()
	if err != nil {
		fmt.Println("wrong", err)
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("invalid task", id)
			continue
		}
		task := tasks[id-1]
		err := mockRemove(task.ID)
		if err != nil {
			fmt.Printf("failed to delete %d.Error %s", id, err)
		} else {
			fmt.Printf("Deleted %d", id)
		}
	}
}

var docmd = &cobra.Command{
	Use:   "do",
	Short: "used to delete task",
	Run:   Delete,
}

func init() {
	RootCmd.AddCommand(docmd)
}
