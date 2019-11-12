package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
	"github.com/shubhanshu7/Gophercises/task/db"
)

var Temp = mockShow
var Temp2 = mockRemove

func TestAdd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	args := []string{"one", "task"}
	a := []string{}
	addCmd.Run(addCmd, args)
	db.Dbcon.Close()
	addCmd.Run(addCmd, a)
}
func TestAllTask(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	arr := []string{"Give", "tasks"}
	listCmd.Run(listCmd, arr)
	db.Dbcon.Close()
	db.Init("sample")
	listCmd.Run(listCmd, arr)

}
func TestDoCmd(t *testing.T) {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "tasks.db")
	db.Init(path)
	validargs := []string{"1", "100", "1989"}
	invalidargs := []string{"10", "hii"}
	docmd.Run(docmd, validargs)
	docmd.Run(docmd, invalidargs)
	defer func() {
		mockShow = Temp
		mockRemove = Temp2
	}()

	mockRemove = func(i int) error {
		return errors.New("My error")
	}
	docmd.Run(docmd, validargs)

	mockShow = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	docmd.Run(docmd, validargs)
}

func TestListFailed(t *testing.T) {

	defer func() {
		mockShow = Temp

	}()

	mockShow = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	s := []string{}
	listCmd.Run(listCmd, s)
}
