package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/shubhanshu7/Gophercises/task/db"
	// "github.com/shubhanshu7/task/db"
)

var Temp = db.AllTasks
var Temp2 = db.DeleteTask

func TestAdd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	args := []string{"Add", "New", "value"}
	a := []string{}
	addCmd.Run(addCmd, args)
	db.Db.Close()
	//store.Init("/")
	addCmd.Run(addCmd, a)
}
func TestList(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	arr := []string{"Hello", "hi"}
	listCmd.Run(listCmd, arr)
	db.Db.Close()
	db.Init("dummy")
	listCmd.Run(listCmd, arr)

}
func TestDoneCmd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	valid_args := []string{"99", "2", "3"}
	invalid_args := []string{"1", "h"}
	docmd.Run(docmd, valid_args)
	docmd.Run(docmd, invalid_args)
	Temp2 = func(i int) error {
		return errors.New("Done")
	}
	docmd.Run(docmd, valid_args)

	Temp = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	docmd.Run(docmd, valid_args)
}

// func TestListNegative(t *testing.T) {

// 	Temp = func() ([]db.Task, error) {
// 		return nil, errors.New("error")
// 	}
// 	s := []string{}
// 	listCmd.Run(listCmd, s)
// }
