package main

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/shubhanshu7/Gophercises/task/cmd"
	"github.com/shubhanshu7/Gophercises/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
