package db

import (
	"errors"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

type teststruct struct {
	input    int
	expected []byte
}

// test boti() and itob()

var hdir, _ = homedir.Dir()

var path = filepath.Join(hdir, "todo.db")
var testval = teststruct{
	input:    5,
	expected: []byte{0, 0, 0, 0, 0, 0, 0, 5},
}

func TestItob(t *testing.T) {
	for index, value := range itob(5) {
		if value != testval.expected[index] {

			t.Error("Itob failed")

		}
	}
}

func TestBtoi(t *testing.T) {
	if btoi(testval.expected) != testval.input {
		t.Error("Btoi failed")
	}
}

func TestInsertTask(t *testing.T) {
	Init(path)
	_, err := CreateTask("Dummy task")
	if err != nil {
		t.Error("Insert Failed")
		fmt.Println(err)
	}
}

func TestRemoveTasks(t *testing.T) {
	err := DeleteTask(5)
	if err != nil {
		t.Error("Delete failed")
	}
}

func TestGetAll(t *testing.T) {
	_, err := AllTasks()
	if err != nil {
		t.Error("Get all failed")
	}
}

func TestInit(t *testing.T) {
	err := Init("/")
	if err != nil {
		errors.New("error")
	}
}

// func TestInitfunc(t *testing.T) {

// 	home, _ := home.Dir()
// 	dbPath := filepath.Join(home, "dummy.db")
// 	err := Init(dbPath)
// 	//err := Init(dbPath)
// 	if err != nil {
// 		t.Error("expected nit got", err)
// 	}
// }
