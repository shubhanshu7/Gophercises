package main

import (
	"log"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/shubhanshu7/Gophercises/task/cmd"
	"github.com/shubhanshu7/Gophercises/task/db"
)

func main() {
	hdir, _ := homedir.Dir()

	path := filepath.Join(hdir, "todo.db")

	HandleError(db.Init(path))

	HandleError(cmd.RootCmd.Execute())
}

// HandleError will check for error
func HandleError(err error) {
	if err != nil {
		log.Println(err.Error())
		return
	}
}
