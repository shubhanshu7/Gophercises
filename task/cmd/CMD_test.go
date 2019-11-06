package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/shubhanshu7/Gophercises/task/db"
)

var Temp = MockShow
var Temp2 = MockRemove

func TestAdd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	args := []string{"one", "task"}
	a := []string{}
	addCmd.Run(addCmd, args)
	db.DbCon.Close()
	addCmd.Run(addCmd, a)
}
func TestAllTask(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	arr := []string{"Give", "tasks"}
	listCmd.Run(listCmd, arr)
	db.DbCon.Close()
	db.Init("sample")
	listCmd.Run(listCmd, arr)

}
func TestDoCmd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	valid_args := []string{"1", "100", "1989"}
	invalid_args := []string{"10", "hii"}
	docmd.Run(docmd, valid_args)
	docmd.Run(docmd, invalid_args)
	defer func() {
		MockShow = Temp
		MockRemove = Temp2
	}()

	MockRemove = func(i int) error {
		return errors.New("My error")
	}
	docmd.Run(docmd, valid_args)

	MockShow = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	docmd.Run(docmd, valid_args)
}

func TestListFailed(t *testing.T) {

	defer func() {
		MockShow = Temp

	}()

	MockShow = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	s := []string{}
	listCmd.Run(listCmd, s)
}
